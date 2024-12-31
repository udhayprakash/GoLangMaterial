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

func increment(name string) {
	// Done() function used to tell that it is done.
	defer waitgroup.Done()

	for _, ch := range name {
		// c++
		// Atomic Functions for fix race condition
		atomic.AddInt32(&c, 1)
		fmt.Printf("%2d) name - %s - char - %q\n", c, name, ch)

		// enter thread in the line by line
		runtime.Gosched()
	}

}

func main() {
	waitgroup.Add(1)
	go increment("one")

	waitgroup.Add(1)
	go increment("two")

	waitgroup.Add(1)
	go increment("jhajsdgasdasdasdjgadgjasjdgasgd")

	// waiting for completion of goroutines.
	waitgroup.Wait()

	// print the counter
	fmt.Println("Counter:", c)
}


// Assignment - try all data types as shared variables, and check for race condition changes
