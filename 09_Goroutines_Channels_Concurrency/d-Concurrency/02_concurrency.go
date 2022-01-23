package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go sendMessage(channel)
	go ReceiveMessage(channel)

	fmt.Scanln()
}

// goroutine that sends the message
func sendMessage(channel chan string) {
	for {
		channel <- "sensing message @" + time.Now().String()
		time.Sleep(5 * time.Second)
	}
}

// goroutine that receives the message
func ReceiveMessage(channel chan string) {
	for {
		message := <-channel
		fmt.Println("Received Message:", message)
	}
}

/*

sending  =======Channel========> Receiving

*/