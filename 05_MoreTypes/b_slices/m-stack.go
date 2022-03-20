package main

import "fmt"

func main() { // stack - LIFO (Last In First Out)
	var stack []string

	// Pushing  - To add elements
	stack = append(stack, "one")
	stack = append(stack, "two")
	stack = append(stack, "three")

	// Peek/Display
	fmt.Println("stack = ", stack)

	// pop - To remove elements
	fmt.Println(stack[len(stack)-1]) // three
	stack = stack[:len(stack)-1]
	fmt.Println("stack = ", stack)

	for _, value := range []string{"four", "five", "six"} {
		stack = append(stack, value)
	}
	fmt.Println("stack = ", stack)

	stack = append(stack, "seven", "eight")
	stack = append(stack, []string{"nine", "ten"}...)
	fmt.Println("stack = ", stack)
	fmt.Println()

	// pop all the elements
	for len(stack) > 0 {
		// Retrieving the last element
		lastPos := len(stack) - 1
		fmt.Println("popped element:", stack[lastPos])

		// popping that last element
		stack = stack[:lastPos]
		// fmt.Println("stack = ", stack)

	}
	fmt.Println("stack = ", stack)

}
