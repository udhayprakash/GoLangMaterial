package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func pinger(c1 chan string) {
	fmt.Println("In pinger function")
	for i := 0; i < 10; i++ {
		fmt.Println("In pinger - sending:")
		// c1 <- "ping"
		c1 <- fmt.Sprintf("\tping %d", i)
	}
}

func printer(c2 chan string) {
	fmt.Println("In printer function")
	for {
		fmt.Println("In printer - receiving:")
		msg := <-c2 // retrieving value from channel
		fmt.Println("\t", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// direct function call
	hello("Golang")

	// goroutine call
	go hello("Golang_Goroutine")

	// var c chan string = make(chan string)
	mychannel := make(chan string)
	go pinger(mychannel)  // source
	go printer(mychannel) // destination

	// work-around to pause the main function till execution
	var input string
	fmt.Scanln(&input)

}
