package main

import (
	"fmt"
)

/*
Purpose: Currying
	- ability to take a function that accepts n parameters and
      turns it into a composition of n functions each of them
      take 1 parameter.
	- Ex: Partial Functions
			- if you have a function that accepts n parameters
              then you can generate from it one of more functions
              with some parameter values already filled in.
*/
func plus(x, y int) int {
	return x + y
}

func partialPlus(x int) func(int) int {
	return func(y int) int {
		return plus(x, y)
	}
}

func main() {
	plusOne := partialPlus(1)
	fmt.Println("plusOne(5):", plusOne(5)) //prints 6
}
