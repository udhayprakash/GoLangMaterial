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
		close(naturals)
	}()

	// Squarer
	go func() {

		// for {
		// 	val := <-naturals    // retrieving value from channel
		// 	squares <- val * val // sending value to "squares" channel
		// }

		for val := range naturals {
			squares <- val * val // sending value to "squares" channel
		}
		close(squares)
	}()

	// // Printer( in main routine)
	// for {
	// 	fmt.Println("<-squares :", <-squares)
	// }

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

}
