package main

import "fmt"

// NewCounter returns a function Count.
// Count prints the number of times it has been invoked.
func NewCounter() (count func()) {
	n := 0
	return func() {
		n++
		fmt.Println("n:", n)
	}

}

func main() {
	counter := NewCounter()

	counter() // n:1
	counter() // n:2
	counter() // n:3

	// fmt.Println(n) // undefined: n

	otherCounter := NewCounter()
	otherCounter() // n : 1

	counter()      // n:4
	otherCounter() // n : 2
}
