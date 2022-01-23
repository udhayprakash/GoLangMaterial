package main

import "fmt"

func main() { // queue - FIFO
	var queue []string

	// Enqueuing
	queue = append(queue, "One")
	queue = append(queue, "Two")
	queue = append(queue, "Three")

	for len(queue) > 0 {
		// retrieving first element
		fmt.Println("queue[0]=", queue[0])

		// Dequeuing
		queue = queue[1:]
	}
}
