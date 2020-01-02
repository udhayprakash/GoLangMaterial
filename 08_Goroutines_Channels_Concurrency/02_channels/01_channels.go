package main

/*
Purpose: Channels
	- are the pipes that connect concurrent goroutines.
	- You can send values into channels from one goroutine
      and receive those values into another goroutine.

*/
import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "google.com" }()
	msg := <-messages
	fmt.Println("Message in Channel:", msg)

	go work(messages)
	msg = <-messages
	fmt.Println( "Message in Channel:", msg)
}

func work(messages chan<- string) {
	messages <- "facebook.com"
}