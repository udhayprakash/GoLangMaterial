package main

import (
	"fmt"
)

// The main function, calling the appleOrOrange() function that will create a race condition
func main() {
	fmt.Printf("Chosen fruit: %s", appleOrOrange())
}

func appleOrOrange() string {

	// Run 'go run -race race.go' to see that this
	// program indeed can create two race conditions,
	// one for each anonymous function

	// This is the variable that is communicated (used in both goroutines)
	// Since each goroutine is scheduled separately, and the interleaving
	// (the sequence in which the instructions in the goroutines are executed)
	// is nondeterministic, i.e., it can change from one run to the next,
	// race conditions can occur when trying to read and write the value
	// of this fruit variable
	var fruit string

	// Anonymous function #1 as a goroutine
	// This sets the value of fruit to 'apple'
	// No control is passed to this goroutine so it doesn't
	// actually get executed, so fruit is not actually
	// assigned a value, but it demonstrates how a race
	// condition can occur
	go func(newFruit string) {
		fruit = newFruit
	}("apple")

	// The same as above, but now passing a different value
	// to another anonymous function that operates on
	// the same fruit variable
	go func(newFruit string) {
		fruit = newFruit
	}("orange")

	return fruit
}