package main

import "fmt"

// NewCounter returns a function Count.
// Count prints the number of times it has been invoked.
func NewCounter() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func main() {
	counter := NewCounter()
	otherCounter := NewCounter()

	counter()      // 1
	counter()      // 2
	counter()      // 3
	otherCounter() // 1 (different n)
	otherCounter() // 2
	counter()      // 4
}
