package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
	pb "server/src/proto"
)

func handleConnection(conn net.Conn) {
	connID := GenerateConnID(conn)
	defer conn.Close()
	defer GlobalManager.DeletePlayer(connID)

	player := GlobalManager.GetOrCreatePlayer(connID)

	// 启动一个协程处理玩家发送的消息
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				log.Println("Connection closed:", err)
				player.QuitChan <- true
				return
			}

			// 解析消息
			var msg pb.Message
			if err := proto.Unmarshal(buffer[:n], &msg); err != nil {
				log.Println("Invalid message:", err)
				continue
			}

			// 将消息发送到玩家的接收通道
			player.RecvChan <- &msg
		}
	}()

	// 启动一个协程处理玩家的响应消息
	go func() {
		for {
			select {
			case rspMsg := <-player.SendChan:
				data, err := proto.Marshal(rspMsg)
				if err != nil {
					log.Println("Failed to marshal response:", err)
					continue
				}
				if _, err := conn.Write(data); err != nil {
					log.Println("Failed to write response:", err)
					player.QuitChan <- true
					return
				}
			case <-player.QuitChan:
				return
			}
		}
	}()
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
