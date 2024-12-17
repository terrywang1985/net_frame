package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
	pb "server/src/proto"
)

// 服务器监听和消息分发
func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		// 解析消息
		var msg pb.Message
		if err := proto.Unmarshal(buffer[:n], &msg); err != nil {
			log.Println("Invalid message:", err)
			continue
		}

		// 临时示例玩家ID（实际应基于连接创建玩家）
		player := GlobalManager.GetOrCreatePlayer(msg.Player)
		player.MsgChan <- &msg
		player.
	}
}

func main() {
	// 初始化消息处理器
	InitMessageHandlers()

	// 启动服务器
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()
	fmt.Println("Server started at :12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
