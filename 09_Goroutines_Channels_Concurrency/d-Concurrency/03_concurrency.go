package main

import (
	"fmt"
	"time"
)

func main() {
	// channel to send message from sender to receiver
	msgChnl := make(chan string)

	// channel to acknowledge message receipt by receiver
	ackChnl := make(chan string)

	go sendMessage2(msgChnl, ackChnl)
	go ReceiveMessage2(msgChnl, ackChnl)

	fmt.Scanln()
}

// goroutine that sends the message
func sendMessage2(msgChannel chan string, ackChannel chan string) {
	for {
		msgChannel <- "sensing message @" + time.Now().String()
		time.Sleep(5 * time.Second)
		ack := <-ackChannel
		fmt.Println("Acknowledgement:", ack)
	}
}

// goroutine that receives the message
func ReceiveMessage2(msgChannel chan string, ackChannel chan string) {
	for {
		message := <-msgChannel
		fmt.Println("Received Message:", message)
		ackChannel <- "message receiver @" + time.Now().String()
	}
}

/*

sending  =======Channel========> Receiving
sendMessage2					ReceiveMessage2
Goroutine						Goroutine
*/
