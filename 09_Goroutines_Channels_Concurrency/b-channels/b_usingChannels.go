package main

import "fmt"

/*
Purpose: Channels
	data := <- myChannel // read from channel myChannel
	myChannel <- data    // write to channel myChannel

*/

/*  ------------------- CHANNELS ARE DESIGNED TO WORK WITH GOROUTINES
func main() {
	messages := make(chan string)

	messages <- "Hello world"

	msg := <-messages
	fmt.Println("\nMessage in Channel:", msg)

}
fatal error: all goroutines are asleep - deadlock!*/

func main() {
	messages := make(chan string)
	fmt.Printf(`
		messages: 
			value: %v
			type: %T
	`, messages, messages)

	go func() {
		// sending message to channel, within function
		messages <- "google.com"
	}()

	// receiving message from channel
	msg := <-messages
	fmt.Println("\nMessage in Channel:", msg)

	go work(messages)

	// receiving message from channel
	msg = <-messages
	fmt.Println("Message in Channel:", msg)

}

func work(messages chan<- string) {
	// sending message to channel, within function
	messages <- "facebook.com"
}
