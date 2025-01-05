// Traditional solution
package main

import (
	"fmt"
	"time"
)

func task(id int, stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Printf("Task %d stopped\n", id)
			return
		default:
			fmt.Printf("Task %d running...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	for i := 1; i <= 3; i++ {
		go task(i, stopChan)
	}
	time.Sleep(2 * time.Second)
	close(stopChan) // Stop all tasks
	time.Sleep(1 * time.Second)
}
