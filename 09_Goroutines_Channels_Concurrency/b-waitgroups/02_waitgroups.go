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
	fmt.Printf("Worder:%d - Started\n", id)

	time.Sleep(time.Second)
	fmt.Printf("\tWorder:%d - Finished\n", id)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
