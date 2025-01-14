package main

import (
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	pb "server/src/proto"
	"sync"
)

//player 与room 之间的映射关系

// Player 玩家结构体
type Player struct {
	Id       string
	Name     string
	Position *pb.Position
	Room     *Room
	Conn     net.Conn
	RecvChan chan *pb.Message // 玩家收消息管道
	SendChan chan *pb.Message // 玩家发消息管道
	QuitChan chan bool        // 退出信号
}

// NewPlayer 创建玩家
func NewPlayer(id string, conn net.Conn) *Player {

	return &Player{
		Id:       id,
		Name:     "",
		Position: &pb.Position{X: 0, Y: 0, Z: 0},

		Conn: conn,

		RecvChan: make(chan *pb.Message, 1000),
		SendChan: make(chan *pb.Message, 1000),
		QuitChan: make(chan bool),
	}
}

// 启动玩家逻辑协程
func (p *Player) Run() {
	var wg sync.WaitGroup
	wg.Add(3) // We have three goroutines to wait for

	defer func() {
		// Clean up when the player exits
		leaveRoom := &Event{
			Type:     EventLeaveRoom,
			PlayerId: p.Id,
		}
		if p.Room != nil {
			p.Room.EventChan <- leaveRoom

			// Wait for the room to process the leave event
			<-leaveRoom.ResponseChan
		}
		close(p.RecvChan)
		close(p.SendChan)
		close(p.QuitChan)
		defer p.Conn.Close()
		defer GlobalManager.DeletePlayer(p.Id)

		log.Printf("Player %s exited", p.Id)
	}()

	// Goroutine to handle incoming messages
	go func() {
		defer wg.Done()
		buffer := make([]byte, 0, 4096) // Initial buffer
		for {
			log.Println("Waiting to read from connection...")

			// Temporary buffer
			tempBuf := make([]byte, 1024)
			n, err := p.Conn.Read(tempBuf)
			if err != nil {
				log.Println("Connection closed:", err)
				p.QuitChan <- true
				return
			}

			// Append read data to buffer
			buffer = append(buffer, tempBuf[:n]...)

			// Process data in buffer
			for {
				log.Printf("Buffer length: %d, Buffer content: %v", len(buffer), buffer)
				// Check if there is enough data to read the packet length
				if len(buffer) < 4 {
					log.Println("Not enough data to read the packet length, breaking out of inner loop")
					break
				}

				// Read packet length
				length := int(binary.LittleEndian.Uint32(buffer[:4]))

				// Check if there is enough data to read the full packet
				if len(buffer) < 4+length {
					break // Not enough data for the full packet
				}

				// Read the full packet
				messageBuf := buffer[4 : 4+length]
				buffer = buffer[4+length:] // Remove processed data

				var parsedMsg pb.Message
				if err := proto.Unmarshal(messageBuf, &parsedMsg); err != nil {
					log.Println("Invalid message:", err)
					continue
				}

				// Attempt to send the message to RecvChan
				select {
				case p.RecvChan <- &parsedMsg:
					// Successfully enqueued
				default:
					// Drop the message if the channel is full
					log.Println("RecvChan full, dropping message")
				}
			}
		}
	}()

	// 处理发送消息的协程
	go func() {
		defer wg.Done()
		for {
			select {
			case rspMsg := <-p.SendChan:
				data, err := proto.Marshal(rspMsg)
				if err != nil {
					log.Println("Failed to marshal response:", err)
					continue
				}

				length := make([]byte, 4)
				binary.LittleEndian.PutUint32(length, uint32(len(data)))

				packet := append(length, data...)
				if _, err := p.Conn.Write(packet); err != nil {
					log.Println("Failed to write response:", err)
					p.QuitChan <- true
					return
				}
			case <-p.QuitChan:
				return
			}
		}
	}()

	// Goroutine to process messages from RecvChan
	go func() {
		defer wg.Done()
		for {
			select {
			case <-p.QuitChan:
				return
			case msg := <-p.RecvChan:

				log.Printf("Received message: %v", msg)
				// Process the message (e.g., handle requests)
				MsgHandler.PlayerHandle(p, msg)
			}
		}
	}()

	wg.Wait() // Wait for all goroutines to finish
}

func (p *Player) SendMessage(msg *pb.Message) {
	p.SendChan <- msg
}

func (p *Player) SendResponse(srcMsg *pb.Message, responseData []byte) {

	// 响应
	response := &pb.Message{
		Id:          srcMsg.GetId() + 1,      // Response ID is request ID + 1
		MsgSerialNo: srcMsg.GetMsgSerialNo(), // Use the same message serial number
		ClientId:    srcMsg.GetClientId(),    // Use the same client ID
		Data:        responseData,
	}

	log.Printf("SendResponse: src: %v, rsp: %v ", srcMsg, response)
	p.SendMessage(response)
}

// HandleMoveRequest 处理移动请求
func (p *Player) HandleMoveRequest(msg *pb.Message) {
	var req pb.MoveRequest
	if err := proto.Unmarshal(msg.GetData(), &req); err != nil {
		log.Println("Failed to parse MoveRequest:", err)
		return
	}
	p.Position = req.Position
	log.Printf("Player %s moved to: %+v", p.Name, p.Position)

	moveEvent := &Event{
		Type:     EventMove,
		PlayerId: p.Id,
		Payload:  req,
	}
	p.Room.EventChan <- moveEvent
}

func (p *Player) HandleLoginRequest(msg *pb.Message) {
	var req pb.LoginRequest
	if err := proto.Unmarshal(msg.GetData(), &req); err != nil {
		log.Println("Failed to parse Login Request:", err)
		return
	}

	p.SendResponse(msg, mustMarshal(&pb.LoginResponse{
		PlayerId: p.Id,
	}))
}

// HandleMoveRequest 处理移动请求
func (p *Player) HandleJoinRoomRequest(msg *pb.Message) {
	var req pb.JoinRoomRequest
	if err := proto.Unmarshal(msg.GetData(), &req); err != nil {
		log.Println("Failed to parse JoinRoomReq:", err)
		return
	}

	result := pb.ErrorCode_OK

	if p.Room != nil {
		log.Printf("Player %s already in room %s", p.Name, p.Room.Name)
		result = pb.ErrorCode_PLAYER_ALREADY_IN_ROOM
		goto sendResponse
	} else {
		room, ok := GlobalManager.GetRoom(req.RoomId)
		if !ok {
			log.Printf("Room %s not found", req.RoomId)
			result = pb.ErrorCode_ROOM_NOT_FOUND
			goto sendResponse
		}
		joinRoomEvent := &Event{
			Type:     EventJoinRoom,
			PlayerId: p.Id,
			Payload:  req,
		}
		room.EventChan <- joinRoomEvent

		response := <-joinRoomEvent.ResponseChan
		if response.(*pb.JoinRoomResponse).Ret == pb.ErrorCode_OK {
			p.Room = room
		}
		log.Printf("Player %s joined room: %s , ret: %d ", p.Name, room.Name, response.(*pb.JoinRoomResponse).Ret)
		p.SendResponse(msg, mustMarshal(response.(*pb.JoinRoomResponse)))
		return
	}

sendResponse:
	p.SendResponse(msg, mustMarshal(&pb.JoinRoomResponse{
		Ret:  result,
		Room: &pb.Room{Id: p.Room.ID, Name: p.Room.Name},
	}))
}

func (p *Player) HandleGetRoomListRequest(msg *pb.Message) {

	var req pb.GetRoomListRequest
	if err := proto.Unmarshal(msg.GetData(), &req); err != nil {
		log.Println("Failed to parse GetRoomListRequest:", err)
		return
	}

	p.SendResponse(msg, mustMarshal(&pb.GetRoomListResponse{
		Ret:   0,
		Rooms: RoomsToProto(GlobalManager.GetAllRooms()),
	}))

}

func (p *Player) HandleCreateRoomRequest(msg *pb.Message) {
	var req pb.CreateRoomRequest
	if err := proto.Unmarshal(msg.GetData(), &req); err != nil {
		log.Println("Failed to parse CreateRoomRequest:", err)
		return
	}

	if p.Room != nil {
		log.Printf("Player %s already in room %s", p.Name, p.Room.Name)
		return
	}

	room := GlobalManager.GetOrCreateRoom(IncrementAndGetRoomCounter(), req.Name)
	room.AddPlayer(p)

	p.SendResponse(msg, mustMarshal(&pb.CreateRoomResponse{
		Ret:  0,
		Room: room.FillRoomMsg(),
	}))

}
