package main

import "fmt"

/*
Purpose: Recursive Functions
http://en.wikipedia.org/wiki/Recursion_(computer_science)

Tail call Optimization (TCO)
	- Avoids creating a new stack when the last call in a recursion
      is the function itself.
	- TCO is not implemented in Golang
*/

// factorial(9) = 9 * 8 * 7 * 6 * .... * 1
func factorial(num int) int {
	fmt.Println("\tnum =", num)
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	result := factorial(9)
	fmt.Println("factorial(9) = ", result)

	result = factorial(7)
	fmt.Println("factorial(7) = ", result)

}

// assignment - what is the max recursion depth in golang

/*

factorial(9) = 9 * factorial(8)
			 = 9 * 8 * factorial(7)

assignment: implement memoization design pattern, using inline functions

	{8: 98789}
*/
