package main

/*
Purpose:
Go’s select lets you wait on multiple channel operations.
select allows multiplexing so we can receive from multiple channels without blocking
Combining goroutines and channels with select is a powerful feature of Go.

*/
import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// msg1 := <-c1
	// msg2 := <-c2
	// fmt.Println("msg1 =", msg1)
	// fmt.Println("msg2 =", msg2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 received", msg2)
		}
	}
	// select - works like a switch but for channels

}
