package main

import "fmt"

// NewCounter returns a function Count.
// Count prints the number of times it has been invoked.
func NewCounter() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println("n:", n)
	}
}

func main() {
	counter := NewCounter()
	otherCounter := NewCounter()

	counter() // n:1
	counter() // n:2
	counter() // n:3

	otherCounter() // n: 1
	otherCounter() // n: 2

	counter() // n:4
}
