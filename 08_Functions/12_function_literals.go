package main

/*
Purpose: Function Literals
	- Also, called as Lamba functions
	- or, Anonymous functions
	- or, inner functions
*/

import "fmt"

func square(x int) func() int {
	return func() int { // anonymous functions
		x = x * x
		return x

	}
}
func main() {
	twoSquared := square(2)
	fmt.Println("2 squared is: ", twoSquared())
}
