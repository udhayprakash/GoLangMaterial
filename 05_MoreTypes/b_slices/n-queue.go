package main

import "fmt"

/*

  ---------------------
<-         QUEUE        <-
  ---------------------

*/

func main() { // queue - FIFO
	var queue []string

	// Enqueuing
	queue = append(queue, "One")
	queue = append(queue, "Two")
	queue = append(queue, "Three")
	fmt.Println(queue)

	for len(queue) > 0 {
		// retrieving first element
		fmt.Println("queue[0]=", queue[0])

		// dequeuing
		queue = queue[1:]
	}
	fmt.Println(queue)
}
