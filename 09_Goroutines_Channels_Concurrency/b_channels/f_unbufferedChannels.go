package main

import (
	"fmt"
)

var myChannel = make(chan int)
var language string

func main() {
	go func() {
		language = "Golang"
		// sending value to channel
		myChannel <- 0

	}()

	// receiving value from channel & discard value
	<-myChannel
	fmt.Println("language:", language)
}

// <-myChannel -- This is blocking code. waiting till some values comes from channel
