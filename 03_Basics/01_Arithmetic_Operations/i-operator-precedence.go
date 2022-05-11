package main

import "fmt"

/*
Purpose: Operator precedence
	 - BODMAS

		B - Brackets
		O - Order of powers or roots
		D - Division
		M - Multiplication
		A - Addition
		S - Subtraction


Unary operators
	- have the highest precedence and bind the strongest.

Binary operators
	------------------------------
	Priority	Operators
	------------------------------
	1			*  /  %  <<  >>  &  &^
	2			+  -  |  ^
	3			==  !=  <  <=  >  >=
	4			&&
	5			||

	Among binary operators multiplication operators bind strongest, followed by addition operators, comparison operators, && (logical and), and finally || (logical or).
	Binary operators of the same precedence associate from left to right:

	x / y * z is the same as (x / y) * z.

Statement operators
	The ++ and -- operators form statements and fall outside the operator hierarchy:
	*p++ is the same as (*p)++.

Examples
	^a >> b is the same as (^a) >> b
	1 + 2*a[i] is the same as 1 + (2*a[i])
	m == n+1 && <-ch > 0 is the same as (m == (n+1)) && ((<-ch) > 0)
*/

func main() {

	// BODMAS
	// Execution order - left to right, and top to bottom
	fmt.Println("1 + 2 -3 =", 1+2-3)

	fmt.Println("1 + 2 -3/3 =", 1+2-3/3)
	// 1 + 2 -3/3
	// 1 + 2 -1
	// 3 -1 = 2

	fmt.Println("1 + 2 - 4 * 3/3 =", 1+2-4*3/3)
	// 1+2-4*3/3
	// 1+2-4*1
	// 1+2-4
	// 3-4 = -1

	fmt.Println("1 +( 2 - 4) * 3/3 =", 1+(2-4)*3/3)
	// 1+(2-4)*3/3
	// 1-2*3/3
	// 1-2*1
	// 1-2 = -1

	// Another Example
	//   12/3*10 - 4
	/*
		4*10 -4
		40 -4
		36


	*/
	fmt.Println(12/3*10 - 4) // 36

	fmt.Println((12/3)*10 - 4)       // 36
	fmt.Println((12 / 3) * (10 - 4)) // 24
	fmt.Println(12/(3*10) - 4)       // -4

}
