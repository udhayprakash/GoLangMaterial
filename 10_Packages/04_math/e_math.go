package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		Pow returns x**y, the base-x exponential of y.
			Pow(x, ±0) = 1 for any x
			Pow(1, y) = 1 for any y
			Pow(x, 1) = x for any x
			Pow(NaN, y) = NaN
			Pow(x, NaN) = NaN
			Pow(±0, y) = ±Inf for y an odd integer < 0
			Pow(±0, -Inf) = +Inf
			Pow(±0, +Inf) = +0
			Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
			Pow(±0, y) = ±0 for y an odd integer > 0
			Pow(±0, y) = +0 for finite y > 0 and not an odd integer
			Pow(-1, ±Inf) = 1
			Pow(x, +Inf) = +Inf for |x| > 1
			Pow(x, -Inf) = +0 for |x| > 1
			Pow(x, +Inf) = +0 for |x| < 1
			Pow(x, -Inf) = +Inf for |x| < 1
			Pow(+Inf, y) = +Inf for y > 0
			Pow(+Inf, y) = +0 for y < 0
			Pow(-Inf, y) = Pow(-0, -y)
			Pow(x, y) = NaN for finite x < 0 and finite non-integer y
	*/
	c3 := math.Pow(2, 3)
	fmt.Printf("%.1f\n", c3)

	/*
		Pow10 returns 10**n, the base-10 exponential of n.
			Pow10(n) =    0 for n < -323
			Pow10(n) = +Inf for n > 308
	*/
	c4 := math.Pow10(2)
	fmt.Printf("%.1f\n", c4)

	/*
		Division
			2)17(8  <--- Quotient
		      16
			------
		       1  <- Remainder
	*/
	fmt.Println("17/2   = ", 17/2)   // 8  ==> Quotient
	fmt.Println("17%2   = ", 17%2)   // 8  ==> Remainder
	fmt.Println("17/2.0 = ", 17/2.0) // 8.5  ===> true result

	/*
		Remainder returns the IEEE 754 floating-point remainder of x/y.
			Remainder(±Inf, y) = NaN
			Remainder(NaN, y) = NaN
			Remainder(x, 0) = NaN
			Remainder(x, ±Inf) = x
			Remainder(x, NaN) = NaN
	*/
	fmt.Println("math.Remainder(10, 2)=", math.Remainder(10, 2))
	fmt.Println("math.Remainder(10, 3)=", math.Remainder(10, 3))
	fmt.Println("math.Remainder(3, 10)=", math.Remainder(3, 10))

	/*
		Mod returns the floating-point remainder of x/y.
		The magnitude of the result is less than y and its sign
		agrees with that of x.
			Mod(±Inf, y) = NaN
			Mod(NaN, y) = NaN
			Mod(x, 0) = NaN
			Mod(x, ±Inf) = x
			Mod(x, NaN) = NaN
	*/
	fmt.Println("\nmath.Mod(10, 2)=", math.Mod(10, 2))
	fmt.Println("math.Mod(10, 3)=", math.Mod(10, 3))
	fmt.Println("math.Mod(3, 10)=", math.Mod(3, 10))

}

// Assignment: math.Remainder vs math.Mod
