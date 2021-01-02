package main

import (
	"fmt"
	"math"
)

func main() {

	// Signbit reports whether x is negative or negative zero.
	fmt.Println("math.Signbit(10)  :", math.Signbit(10))
	fmt.Println("math.Signbit(-10) :", math.Signbit(-10))

	// Copysign returns a value with the magnitude of x and the sign of y.
	// func Copysign(x, y float64) float64
	fmt.Println("math.Copysign(12, 10)  =", math.Copysign(12, 10))
	fmt.Println("math.Copysign(-12, 10) =", math.Copysign(-12, 10))
	fmt.Println("math.Copysign(12, -10) =", math.Copysign(12, -10))

	/* Dim returns the maximum of x-y or 0.
	func Dim(x, y float64) float64
		Dim(+Inf, +Inf) = NaN
		Dim(-Inf, -Inf) = NaN
		Dim(x, NaN) = Dim(NaN, x) = NaN
	*/
	fmt.Println("\nmath.Dim(7, 3) =", math.Dim(7, 3))
	fmt.Println("math.Dim(3, 7) =", math.Dim(3, 7))
	fmt.Println("math.Dim(3, 3) =", math.Dim(3, 3))
	/*
		Max returns the larger of x or y.
			Max(x, +Inf) = Max(+Inf, x) = +Inf
			Max(x, NaN) = Max(NaN, x) = NaN
			Max(+0, ±0) = Max(±0, +0) = +0
			Max(-0, -0) = -0
	*/
	fmt.Println("\nmath.Max(7, 3)=", math.Max(7, 3))
	fmt.Println("math.Max(3, 7)=", math.Max(3, 7))
	fmt.Println("math.Max(3,-7)=", math.Max(3, -7))

	/*
		Min returns the smaller of x or y.
			Min(x, -Inf) = Min(-Inf, x) = -Inf
			Min(x, NaN) = Min(NaN, x) = NaN
			Min(-0, ±0) = Min(±0, -0) = -0
	*/
	fmt.Println("\nmath.Min(7, 3)=", math.Min(7, 3))
	fmt.Println("math.Min(3, 7)=", math.Min(3, 7))
	fmt.Println("math.Min(3,-7)=", math.Min(3, -7))

}
