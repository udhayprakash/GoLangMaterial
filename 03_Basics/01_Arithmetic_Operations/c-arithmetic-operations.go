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

	/*
				3 ) 10 (3   <-- Quotient  /
					 9
					---
		             1  <- remainder  %

	*/
	fmt.Println("10 / 3 = ", 10/3) // Quotient
	fmt.Println("10 % 3 = ", 10%3) // Remainder

	fmt.Println(10 / 3.0) // 3.3333333333333335
	// NOTE: to get true division result, either nmerator or denominator should be float

	fmt.Println(10 / float64(3)) // 3.3333333333333335

	fmt.Println("\n compound operators +=, -=, *=, /=, %=, ++, -- ")

	// 10 += 1 // Error - cannot assign to 10

	num1 := 10

	num1 = num1 + 1
	fmt.Println("num1 = ", num1) // 11

	num1 += 1
	fmt.Println("num1 = ", num1) // 12

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
