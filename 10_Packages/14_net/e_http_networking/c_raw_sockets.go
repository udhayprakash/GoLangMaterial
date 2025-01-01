package main

import (
	"fmt"
	"net"
)

// Work with raw sockets for low-level network communication

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := "Custom TCP Packet"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending packet:", err)
		return
	}
	fmt.Println("Custom packet sent")
}
