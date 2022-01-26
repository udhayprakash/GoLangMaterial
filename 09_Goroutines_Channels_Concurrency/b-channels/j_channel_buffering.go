package main

/*
Purpose: Channel Buffering
	- By default channels are unbuffered, meaning that they
      will only accept sends (chan <-) if there is a
      corresponding receive (<- chan) ready to receive the
      sent value.
	- Buffered channels accept a limited number of values
      without a corresponding receiver for those values.
*/
import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println("\n<-messages :", <-messages)
	fmt.Println("<-messages :", <-messages)

	messages <- "udhay"
	messages <- "prakash"
	// messages <- "prakash2"
	fmt.Println("\n<-messages :", <-messages)
	fmt.Println("<-messages :", <-messages)

	messages <- "udhay"
	fmt.Println("\n<-messages :", <-messages)
}

// NOTE: "fatal error: all goroutines are asleep - deadlock!"
// This error occures
// 1. when there is no value in the channel
// 2. If you trying to add values , more than the channel buffer size
