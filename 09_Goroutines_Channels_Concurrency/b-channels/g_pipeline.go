package main

import "fmt"

func main() {
	// Unbuffered (or SYnchronous) Channels
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x // sending value to channel
		}
	}()

	// Squarer
	go func() {

		for {
			val := <-naturals    // retrieving value from channel
			squares <- val * val // sending value to "squares" channel
		}
	}()

	// Printer( in main routine)
	for {
		fmt.Println("<-squares :", <-squares)
	}

}
