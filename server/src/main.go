package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	connID := GenerateConnID(conn)
	defer GlobalManager.DeletePlayer(connID)

	player := GlobalManager.GetOrCreatePlayer(connID, conn)

	// Wait for the player to signal that it has quit
	<-player.QuitChan
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
