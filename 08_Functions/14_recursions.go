package main

/*
Purpose: Recursive Functions
http://en.wikipedia.org/wiki/Recursion_(computer_science)

Tail call Optimization (TCO)
	- Avoids creating a new stack when the last call in a recursion
      is the function itself.
	- TCO is not implemented in Golang
*/

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	fmt.Println("factorial(9) =", factorial(9))

}

// TODO - fibonacci series generation
