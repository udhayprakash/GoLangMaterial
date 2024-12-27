package main

import "fmt"

func main() { // stack - LIFO (Last In First Out)
	var stack []string

	// stack[0] = "sa"  panic: runtime error: index out of range [0] with length 0

	// Pushing  - To add elements
	stack = append(stack, "one")
	stack = append(stack, "two", "three")

	// Peek/Display
	fmt.Println("stack = ", stack)

	// pop - To remove elements

	// stack -- [one two three ]
	// index  -   0   1    2
	// length  -- 3
	fmt.Println(stack[len(stack)-1]) // three

	stack = stack[:len(stack)-1]
	fmt.Println("stack = ", stack) //  [one two]

	// insert operation
	for _, value := range []string{"four", "five", "six"} {
		stack = append(stack, value)
	}
	fmt.Println("stack = ", stack) // [one two four five six]

	stack = append(stack, "seven", "eight")
	fmt.Println("stack = ", stack) // [one two four five six seven eight]

	// ... spread operator
	stack = append(stack, []string{"nine", "ten"}...)
	fmt.Println("stack = ", stack) // [one two four five six seven eight nine ten]
	fmt.Println()

	// pop all the elements
	for len(stack) > 0 {
		// Retrieving the last element
		lastPos := len(stack) - 1
		fmt.Println("popped element:", stack[lastPos])

		// popping that last element
		stack = stack[:lastPos]
		fmt.Println("\tstack = ", stack)
	}
	fmt.Println("stack = ", stack)

}
