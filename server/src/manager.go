package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type Manager struct {
	rooms   sync.Map // 使用 sync.Map 来存储房间
	players sync.Map // 使用 sync.Map 来存储玩家
}

var roomCounter uint64

func IncrementAndGetRoomCounter() uint64 {
	return atomic.AddUint64(&roomCounter, 1)
}

// 创建房间
func (rm *Manager) GetOrCreateRoom(id uint64, name string) *Room {
	room, loaded := rm.rooms.LoadOrStore(id, NewRoom(id, name)) // 从 sync.Map 获取房间，如果不存在则创建
	// 如果房间是新创建的，则启动其协程
	if !loaded {
		log.Printf("Room created: %s", name)
		go room.(*Room).Run() // 启动房间逻辑协程
	} else {
		log.Printf("Room already exists: %s", name)
	}
	return room.(*Room)
}

// 获取房间
func (rm *Manager) GetRoom(id string) (*Room, bool) {
	room, ok := rm.rooms.Load(id) // 从 sync.Map 获取房间
	if !ok {
		log.Printf("Room %s not found", id)
		return nil, false
	}
	return room.(*Room), true
}

// 删除房间
func (rm *Manager) DeleteRoom(id string) {
	if room, ok := rm.GetRoom(id); ok {
		room.QuitChan <- true // 关闭房间
		rm.rooms.Delete(id)   // 删除房间
		log.Printf("Room %s deleted", id)
	}
}

// 获取所有房间
func (rm *Manager) GetAllRooms() []*Room {
	var rooms []*Room
	rm.rooms.Range(func(key, value interface{}) bool {
		rooms = append(rooms, value.(*Room)) // 将所有房间收集到数组中
		return true                          // 继续遍历
	})
	return rooms
}

// 创建玩家
func (rm *Manager) GetOrCreatePlayer(id string) *Player {
	player, loaded := rm.players.LoadOrStore(id, NewPlayer(id)) // 从 sync.Map 获取玩家，如果不存在则创建

	// 如果玩家是新创建的，则启动其协程
	if !loaded {
		log.Printf("Player created: %s", id)
		go player.(*Player).Run() // 启动玩家逻辑协程
	} else {
		log.Printf("Player already exists: %s", id)
	}

	return player.(*Player)
}

// 获取玩家
func (rm *Manager) GetPlayer(id string) (*Player, bool) {
	player, ok := rm.players.Load(id) // 从 sync.Map 获取玩家
	if !ok {
		log.Printf("Player %s not found", id)
		return nil, false
	}
	return player.(*Player), true
}

// 删除玩家
func (rm *Manager) DeletePlayer(id string) {
	if player, ok := rm.GetPlayer(id); ok {
		player.QuitChan <- true // 退出玩家
		rm.players.Delete(id)   // 删除玩家
		log.Printf("Player %s deleted", id)
	}
}

// 获取所有玩家
func (rm *Manager) GetAllPlayers() []*Player {
	var players []*Player
	rm.players.Range(func(key, value interface{}) bool {
		players = append(players, value.(*Player)) // 将所有玩家收集到数组中
		return true                                // 继续遍历
	})
	return players
}

// 全局管理器实例
var GlobalManager = &Manager{}

func init() {
	// 在程序启动时，如果需要做初始化工作，可以在这里进行
	log.Println("Manager initialized")
}
