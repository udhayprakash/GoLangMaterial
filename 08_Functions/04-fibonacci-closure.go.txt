package main

import (
	"fmt"
)

// Define the number of numbers we want to print
var displayedNumbers int = 10

// Define a struct for fibonacci numbers
type fibStruct struct {
	Last int
	Curr int
}

// fibonacci is a function that returns a function that returns an int
func fibonacci() func() int {
	fib := &fibStruct{0, 0}

	// Returns an anonymous function which returns an int
	return func() int {
		// Check current value by adding the 2 last values
		// We need that temporary variable as we need to update old variables
		// with new values
		newFib := fib.Last + fib.Curr

		// Update n and n-1 values
		// If sum was 0 we increment last, not current so that next time
		// sum is 1 and with normal updates (curr = sum == 1 & last = curr == 0)
		// we ensure the double 1 as is normal fibonacci
		if newFib == 0 {
			fib.Last = 1
		} else {
			fib.Last = fib.Curr
		}

		fib.Curr = newFib

		// Return current value
		return newFib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < displayedNumbers; i++ {
		fmt.Println(f())
	}
}
