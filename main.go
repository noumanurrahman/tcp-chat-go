package main

import (
	"fmt"
	"log"
	"net"

	"tcp-chat/chat"
)

func main() {
	server := chat.NewServer()
	go server.Run()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error starting server %s", err.Error())
	}

	defer listener.Close()
	fmt.Println("Server started on port 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection %s", err.Error())
			continue
		}
		go server.NewClient(conn)
	}
}
