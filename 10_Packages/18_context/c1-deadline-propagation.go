// Traditional solution
// Example: Setting a hard deadline for a task.

package main

import (
	"fmt"
	"time"
)

func processTask(deadline time.Time) {
	if time.Now().After(deadline) {
		fmt.Println("Task cancelled: deadline exceeded")
		return
	}
	fmt.Println("Processing task...")
	time.Sleep(2 * time.Second)
	fmt.Println("Task completed")
}

func main() {
	deadline := time.Now().Add(1 * time.Second)
	go processTask(deadline)
	time.Sleep(3 * time.Second) // Wait to see the output
}
