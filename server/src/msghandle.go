package main

import (
	pb "server/src/proto"
)

// 消息管理器
type MessageManager struct {
	player_handlers map[pb.MessageId]func(player *Player, msg *pb.Message)
	room_handlers   map[pb.MessageId]func(room *Room, msg *pb.Message)
}

// 初始化消息管理器
func NewMessageManager() *MessageManager {
	return &MessageManager{
		player_handlers: make(map[pb.MessageId]func(player *Player, msg *pb.Message)),
		room_handlers:   make(map[pb.MessageId]func(room *Room, msg *pb.Message)),
	}
}

// 注册消息处理回调
func (m *MessageManager) PlayerRegister(msgId pb.MessageId, handler func(player *Player, msg *pb.Message)) {
	m.player_handlers[msgId] = handler
}

// 注册消息处理回调
func (m *MessageManager) RoomRegister(msgId pb.MessageId, handler func(player *Room, msg *pb.Message)) {
	m.room_handlers[msgId] = handler
}

// 处理消息
func (m *MessageManager) PlayerHandle(player *Player, msg *pb.Message) {
	if handler, ok := m.player_handlers[msg.GetId()]; ok {
		handler(player, msg)
	}
}

// 房间处理消息
func (m *MessageManager) RoomHandle(room *Room, msg *pb.Message) {
	if handler, ok := m.room_handlers[msg.GetId()]; ok {
		handler(room, msg)
	}
}

// 全局消息管理器实例
var MsgHandler = NewMessageManager()

// 注册所有消息回调
func InitMessageHandlers() {
	MsgHandler.PlayerRegister(pb.MessageId_GET_ROOM_LIST_REQUEST, (*Player).HandleGetRoomListRequest)
	MsgHandler.PlayerRegister(pb.MessageId_CREATE_ROOM_REQUEST, (*Player).HandleCreateRoomRequest)
	MsgHandler.PlayerRegister(pb.MessageId_MOVE_REQUEST, (*Player).HandleMoveRequest)

	MsgHandler.RoomRegister(pb.MessageId_MOVE_REQUEST, (*Room).HandleMoveRequest)
}
