package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	pb "server/src/proto"
)

//player 与room 之间的映射关系

// Player 玩家结构体
type Player struct {
	ID       string
	Name     string
	Position *pb.Position
	Room     *Room
	RecvChan chan *pb.Message // 玩家收消息管道
	SendChan chan *pb.Message // 玩家发消息管道
	QuitChan chan bool        // 退出信号
}

// NewPlayer 创建玩家
func NewPlayer(id string) *Player {
	return &Player{
		ID:       id,
		Name:     "",
		Position: &pb.Position{X: 0, Y: 0, Z: 0},
		RecvChan: make(chan *pb.Message, 10),
		SendChan: make(chan *pb.Message, 10),
		QuitChan: make(chan bool),
	}
}

// 启动玩家逻辑协程
func (p *Player) Run() {
	for {
		select {
		case msg := <-p.RecvChan:
			// 处理消息
			MsgHandler.PlayerHandle(p, msg)
		case <-p.QuitChan:
			fmt.Printf("Player %s quitting...\n", p.Name)
			return
		}
	}
}

func (p *Player) SendMessage(msg *pb.Message) {
	p.SendChan <- msg
}

// HandleMoveRequest 处理移动请求
func (p *Player) HandleMoveRequest(data []byte) {
	var req pb.MoveRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		log.Println("Failed to parse MoveRequest:", err)
		return
	}
	p.Position = req.Position
	log.Printf("Player %s moved to: %+v", p.Name, p.Position)

	// 同步给房间内其他玩家
	response := &pb.Message{
		Id: pb.MessageId_MOVE_RESPONSE,
		Data: mustMarshal(&pb.MoveResponse{
			Ret:  0,
			Room: &pb.Room{Id: p.Room.ID, Name: p.Room.Name},
		}),
	}

	p.RecvChan <- response
}

func (p *Player) HandleGetRoomListRequest(data []byte) {
	var req pb.GetRoomListRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		log.Println("Failed to parse GetRoomListRequest:", err)
		return
	}

	// 响应
	response := &pb.Message{
		Id: pb.MessageId_GET_ROOM_LIST_RESPONSE,
		Data: mustMarshal(&pb.GetRoomListResponse{
			Ret:   0,
			Rooms: RoomsToProto(GlobalManager.GetAllRooms()),
		}),
	}

	p.SendMessage(response)
}

func (p *Player) HandleCreateRoomRequest(data []byte) {
	var req pb.CreateRoomRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		log.Println("Failed to parse CreateRoomRequest:", err)
		return
	}

	if p.Room != nil {
		log.Printf("Player %s already in room %s", p.Name, p.Room.Name)
		return
	}

	room := GlobalManager.GetOrCreateRoom(IncrementAndGetRoomCounter(), req.Name)
	room.AddPlayer(p)

	// 响应
	response := &pb.Message{
		Id: pb.MessageId_CREATE_ROOM_RESPONSE,
		Data: mustMarshal(&pb.CreateRoomResponse{
			Ret:  0,
			Room: room.FillRoomMsg(),
		}),
	}

	p.SendMessage(response)
}
