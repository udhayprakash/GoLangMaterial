package main

import (
	"fmt"
	"time"
)

// Channels provide a way for two goroutines to communicate
// with each other and synchronize their execution.

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"  // placing the value into channel
	}
}

func printer(c chan string){
	for {
		msg := <-c // retrieving value from channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	// work-around to pause the main function till execution
	var input string
	fmt.Println("Type Enter to close program execution")
	fmt.Scanln(&input)

}
