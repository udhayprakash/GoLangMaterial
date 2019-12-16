package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ //i += 1  // i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// all unary operations (i++, i--, --i, ++i) are not supported
