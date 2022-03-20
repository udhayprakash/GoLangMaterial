package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	// sending to a closed channel
	//queue <- "four" // panic: send on closed channel

	// receiving from a closed channel
	fmt.Println("<-queue =", <-queue)

	for ele := range queue {
		fmt.Println("ele = ", ele)
	}

}
