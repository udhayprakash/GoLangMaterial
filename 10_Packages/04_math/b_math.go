package main

import (
	"fmt"
	"math"
)

func main() {

	/* Absolute results in positive value
	Abs(±Inf) = +Inf
	Abs(NaN) = NaN
	*/
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x) // 2.0

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y) // 2.0

	// Trigonometric functions
	// Acos(x) = NaN if x < -1 or x > 1
	fmt.Printf("%.2f", math.Acos(1))

	/*
		Acosh(+Inf) = +Inf
		Acosh(x) = NaN if x < 1
		Acosh(NaN) = NaN
	*/
	fmt.Printf("%.2f\n", math.Acosh(1))

	/*
		Asin(±0) = ±0
		Asin(x) = NaN if x < -1 or x > 1
	*/
	fmt.Printf("%.2f\n", math.Asin(0))

	/*
		Asinh(±0) = ±0
		Asinh(±Inf) = ±Inf
		Asinh(NaN) = NaN
	*/
	fmt.Printf("%.2f\n", math.Asinh(0))

	/*
		Atan(±0) = ±0
		Atan(±Inf) = ±Pi/2
	*/
	fmt.Printf("%.2f\n", math.Atan(0))

	/*
		Hypothenuse of a right triangle
	*/
}
