package main

type EventType int

const (
	EventJoinRoom EventType = iota
	EventLeaveRoom
	EventChat
	EventMove
)

type Event struct {
	Type         EventType
	PlayerId     string
	Payload      interface{}
	ResponseChan chan interface{} // 用于向玩家协程返回结果
}

// 事件管理器
type EventManager struct {
	event_handlers map[EventType]func(room *Room, event *Event)
}

// 初始化管理器
func NewEventManager() *EventManager {
	return &EventManager{
		event_handlers: make(map[EventType]func(room *Room, event *Event), 100),
		//room_handlers:   make(map[pb.MessageId]func(room *Room, roomMsg *RoomMessage)),
	}
}

// // 注册事件回调
func (em *EventManager) Register(evtType EventType, handler func(room *Room, event *Event)) {
	em.event_handlers[evtType] = handler
}

// 处理消息
func (em *EventManager) Handle(room *Room, event *Event) {
	if handler, ok := em.event_handlers[event.Type]; ok {
		handler(room, event)
	}
}

// 全局消息管理器实例
var EventHandler = NewEventManager()

// 注册所有消息回调
func InitEventHandlers() {
	EventHandler.Register(EventJoinRoom, (*Room).HandleJoinRoom)
	EventHandler.Register(EventLeaveRoom, (*Room).HandleLeaveRoom)
	EventHandler.Register(EventChat, (*Room).HandleChat)
	EventHandler.Register(EventMove, (*Room).HandleMove)
}
