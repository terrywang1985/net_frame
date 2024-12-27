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
	ID       string
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
		ID:       id,
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

	defer p.Conn.Close()
	defer GlobalManager.DeletePlayer(p.ID)

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

				// Attempt to send the message to RecvChan
				select {
				case p.RecvChan <- &pb.Message{Data: messageBuf}:
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
				var parsedMsg pb.Message
				if err := proto.Unmarshal(msg.Data, &parsedMsg); err != nil {
					log.Println("Invalid message:", err)
					continue
				}
				log.Printf("Received message: %v", parsedMsg)
				// Process the message (e.g., handle requests)
				MsgHandler.PlayerHandle(p, &parsedMsg)
				if p.Room != nil {
					MsgHandler.RoomHandle(p.Room, &parsedMsg)
				}
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

	p.SendResponse(msg, mustMarshal(&pb.MoveResponse{
		Ret:  0,
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
