package main

/*
Purpose: Channels
	data := <- myChannel // read from channel myChannel
	myChannel <- data    // write to channel myChannel

*/
import "fmt"

func main() {
	messages := make(chan string)
	fmt.Printf(`
		messages: 
			value: %v
			type: %T`, messages, messages)


	go func() { messages <- "google.com" }()
	msg := <-messages
	fmt.Println("\nMessage in Channel:", msg)


	go work(messages)
	msg = <-messages
	fmt.Println( "Message in Channel:", msg)

}

func work(messages chan<- string) {
	messages <- "facebook.com"
}