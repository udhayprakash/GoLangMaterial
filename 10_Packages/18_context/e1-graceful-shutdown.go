// Traditional solution
package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker stopping...")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go worker(stopChan)
	time.Sleep(2 * time.Second)
	close(stopChan) // Signal worker to stop
	time.Sleep(1 * time.Second)
}
