package main

import (
	"fmt"
)

func main() { // stack - LIFO
	var stack []string

	// Pushing
	stack = append(stack, "One")
	stack = append(stack, "Two")
	stack = append(stack, "Three")

	// Display
	fmt.Println("stack =", stack)

	for len(stack) > 0 {
		// Retrieving top element
		n := len(stack) - 1
		fmt.Println(stack[n])

		// Poping from stack
		stack = stack[:n]
	}
}
