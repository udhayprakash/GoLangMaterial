/* Purpose: Golang program to fix the race
condition using atomic package
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// All goroutines will increment variable c
// waitgroup is waiting for the completion
// of program.
var (
	c         int32
	waitgroup sync.WaitGroup
)

func main() {

	// with the help of Add() function add one
	// for each goroutine a count of total 3
	waitgroup.Add(3)

	// increment with the help
	// of increment() function
	go increment("one")
	go increment("two")
	go increment("three")

	// waiting for completion of goroutines.
	waitgroup.Wait()

	// print the counter
	fmt.Println("Counter:", c)

}

func increment(name string) {

	// Done() function used to tell that it is done.
	defer waitgroup.Done()

	for range name {

		// Atomic Functions for fix race condition
		atomic.AddInt32(&c, 1)

		// enter thread in the line by line
		runtime.Gosched()
	}
}
