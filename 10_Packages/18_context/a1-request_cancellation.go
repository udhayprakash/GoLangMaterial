// Traditionally request cancellation, without context
package main

import (
	"fmt"
	"time"
)

func longRunningOperation(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Operation cancelled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go longRunningOperation(stopChan)
	time.Sleep(2 * time.Second)
	close(stopChan) // Cancel the operation
	time.Sleep(1 * time.Second)
}
