package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start Script")

	for i := 0; i < 10; i++ {
		go printItem(i)
	}

	fmt.Println("End Script")

	// Wait, giving time for the go routines
	// to finish.
	time.Sleep(1000)
}

func printItem(i int) {
	fmt.Printf("Print Item: %v\n", (i + 1))
}