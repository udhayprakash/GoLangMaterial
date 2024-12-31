package main

import (
	"fmt"
	"sync"
	"time"
)

/*
To wait for multiple goroutines to finish, we can use a wait group.

sync.WaitGroup is struct type with three methods:
	- Add: It adds an integer counter number, that says main function
		   will wait for that number of goroutines to complete.
	- Done: It decreases the count, which is added in Add() method
		- It should be called at the end of goroutine
	- Wait: It waits for all goroutines to finish.
		- When count reaches 0, it finishes the wait.
*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}
	
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers done")
}
