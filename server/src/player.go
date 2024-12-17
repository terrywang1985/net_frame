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
	MsgChan  chan *pb.Message // 玩家消息管道
	QuitChan chan bool        // 退出信号
}

// NewPlayer 创建玩家
func NewPlayer(id string) *Player {
	return &Player{
		ID:       id,
		Name:     "",
		Position: &pb.Position{X: 0, Y: 0, Z: 0},
		MsgChan:  make(chan *pb.Message, 10),
		QuitChan: make(chan bool),
	}
}

// 启动玩家逻辑协程
func (p *Player) Run() {
	for {
		select {
		case msg := <-p.MsgChan:
			// 处理消息
			MsgHandler.PlayerHandle(p, msg)
		case <-p.QuitChan:
			fmt.Printf("Player %s quitting...\n", p.Name)
			return
		}
	}
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
	p.Room.Broadcast(p.ID, response)
}
