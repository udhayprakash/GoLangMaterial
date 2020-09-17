package main

import "fmt"

/*
Purpose: Operator precedence

Unary operators
	- have the highest precedence and bind the strongest.

Binary operators
	------------------------------
	Prio	Operators
	------------------------------
	1		*  /  %  <<  >>  &  &^
	2		+  -  |  ^
	3		==  !=  <  <=  >  >=
	4		&&
	5		||

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
	x := 10
	y := 20
	z := 30

	fmt.Println("x / y * z", x/y*z)
}
