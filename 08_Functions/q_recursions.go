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

func factorial(num int) int {
	// 9 -> 9 * 8 * 7 * 6 * .... * 1
	if num <= 1 {
		return 1
	}
	return num * factorial(num -1)
}


func main(){
	result := factorial(9)
	fmt.Println("factorial(9) = ", result)

	fmt.Println("factorial(6):", factorial(6))

	for i:= 0; i < 12; i++ {
		fmt.Printf("factorial(%2d)=%8d\n", i, factorial(i))
	}
}