package main

import (
	"log"

	"google.golang.org/protobuf/proto"
	pb "server/src/proto"
	"sync"
)

type Room struct {
	ID       uint64
	Name     string
	Players  map[string]*Player
	MsgChan  chan *RoomMessage // 房间消息管道
	QuitChan chan bool         // 退出信号
	Mutex    sync.Mutex        // 保护 Players
}

type RoomMessage struct {
	PlayerID string
	Message  *pb.Message
}

// 创建一个房间
func NewRoom(id uint64, name string) *Room {
	return &Room{
		ID:       id,
		Name:     name,
		Players:  make(map[string]*Player),
		MsgChan:  make(chan *RoomMessage, 100),
		QuitChan: make(chan bool),
	}
}

// 启动房间协程
func (r *Room) Run() {
	log.Printf("Room %s is running...\n", r.Name)
	for {
		select {
		case roomMsg := <-r.MsgChan:
			r.HandleMessage(roomMsg)
		case <-r.QuitChan:
			log.Printf("Room %s is closing...", r.Name)
			return
		}
	}
}

// 处理房间内的消息
func (r *Room) HandleMessage(msg *RoomMessage) {
	switch msg.Message.Id {
	case pb.MessageId_MOVE_REQUEST:
		// 解析消息
		var req pb.MoveRequest
		if err := proto.Unmarshal(msg.Message.Data, &req); err != nil {
			log.Println("Failed to parse MoveRequest:", err)
			return
		}

		// 更新玩家位置
		player := r.Players[msg.PlayerID]
		player.Position = req.Position
		log.Printf("Player %s moved to %+v", player.Name, player.Position)

		// 广播给其他玩家
		response := &pb.Message{
			Id: pb.MessageId_MOVE_RESPONSE,
			Data: mustMarshal(&pb.MoveResponse{
				Ret:  0,
				Room: &pb.Room{Id: r.ID, Name: r.Name},
			}),
		}
		r.Broadcast(msg.PlayerID, response)
	}
}

// HandleMoveRequest 处理移动请求
func (r *Room) HandleMoveRequest(data []byte) {
	var req pb.MoveRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		log.Println("Failed to parse MoveRequest:", err)
		return
	}

	// 更新玩家位置
	player := r.Players[req.PlayerId]
	player.Position = req.Position
	log.Printf("Player %s moved to %+v", player.Name, player.Position)

	noti := &pb.Message{
		Id: pb.MessageId_ROOM_STATE_NOTIFICATION,
		Data: mustMarshal(&pb.RoomStateNotification{
			Room: r.FillRoomMsg(),
		}),
	}

	r.Broadcast(req.PlayerId, noti)
}

func (r *Room) FillRoomMsg() *pb.Room {
	room := &pb.Room{
		Id:   r.ID,
		Name: r.Name,
	}
	room.Players = make([]*pb.Player, 0)
	for _, player := range r.Players {
		room.Players = append(room.Players, &pb.Player{
			Id:       player.ID,
			Name:     player.Name,
			Position: player.Position,
		})
	}
	return room
}

// 广播消息给所有玩家（排除发送者）
func (r *Room) Broadcast(excludePlayerID string, msg *pb.Message) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	for id, player := range r.Players {
		if id != excludePlayerID {
			player.SendMessage(msg)
		}
	}
}

// 添加玩家
func (r *Room) AddPlayer(player *Player) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Players[player.ID] = player
	player.Room = r
	log.Printf("Player %s joined room %s", player.Name, r.Name)
}

// mustMarshal marshals a protobuf message and logs a fatal error if it fails.
func mustMarshal(pb proto.Message) []byte {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalf("Failed to marshal protobuf message: %v", err)
	}
	return data
}
