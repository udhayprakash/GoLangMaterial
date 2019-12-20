package main

/*
Purpose: Recursive Functions
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
