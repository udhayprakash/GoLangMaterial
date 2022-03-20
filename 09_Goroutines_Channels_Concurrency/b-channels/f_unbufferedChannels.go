package main

import (
	"fmt"
)

var language string
var myChannel = make(chan int)

func main() {
	go func() {
		language = "Golang"
		// sending value to channel
		myChannel <- 0
	}()

	// NO more time.sleep or fmt.Scanf

	// receiving value from channel & discard value
	<-myChannel

	fmt.Println("language:", language)
}

// <-myChannel -- This is blocking code.
// waiting till some values comes from channel
