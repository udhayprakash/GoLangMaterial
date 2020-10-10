package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "from 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "from 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("Done")
}

/*
Output:

received from 1
received from 2
Done
*/
