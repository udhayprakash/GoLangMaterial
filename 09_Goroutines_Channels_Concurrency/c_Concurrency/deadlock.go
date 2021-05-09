package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Problem
	// c <- 1
	// // fatal error: all goroutines are asleep - deadlock!

	// Solution
	go func() {
		c <- 1
	}()

	fmt.Println("<-c", <-c)

}
