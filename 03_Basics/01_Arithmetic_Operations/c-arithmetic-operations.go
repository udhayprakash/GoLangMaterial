package main

import (
	"fmt"
)

/*
arithmetic operations
	+ 	addition
	-	subtraction
	*	Multiplication

	/	Division
	%	Remainder

*/

func main() {
	fmt.Println("10 + 3 = ", 10+3)
	fmt.Println("10 - 3 = ", 10-3)
	fmt.Println("10 * 3 = ", 10*3)

	fmt.Println("10 / 3 = ", 10/3) // Quotient
	fmt.Println("10 % 3 = ", 10%3) // Remainder

	/*
				3 ) 10 (3   <-- Quotient  /
					 9
					---
		             1  <- remainder  %

	*/
	fmt.Println()

	// 	compound operators +=, -=, *=, /=, %=, ++, --
	// 10 += 1 // Error - cannot assign to 10

	num1 := 10

	num1 = num1 + 1
	fmt.Println("num1 = ", num1)

	num1 += 1
	fmt.Println("num1 = ", num1)

	num1++ // num1 += 1
	fmt.Println("num1 = ", num1)

	num1 -= 1 // num1 = num1 - 1
	fmt.Println("num1 = ", num1)

	num1-- // num1 -= 1
	fmt.Println("num1 = ", num1)

	num1 *= 1 // num1 = num1 * 1
	fmt.Println("num1 = ", num1)

	num1 /= 1 // num1 = num1 / 1
	fmt.Println("num1 = ", num1)

	num1 %= 1 // num1 = num1 % 1
	fmt.Println("num1 = ", num1)

}
