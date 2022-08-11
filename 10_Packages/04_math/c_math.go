package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("math.Exp(1)  = ", math.Exp(1))
	fmt.Println("math.Exp(10) = ", math.Exp(10))
	for x := 0; x < 8; x++ {
		fmt.Printf("e^%d = %8.3f\n", x, math.Exp(float64(x)))
	}

	/*
		Exp returns e**x, the base-e exponential of x.
		func Exp(x float64) float64
			Exp(+Inf) = +Inf
			Exp(NaN) = NaN
	*/
	fmt.Println("math.Exp(10)     =", math.Exp(10))
	fmt.Println("math.Exp(math.E) =", math.Exp(math.E))

	// Exp2 returns 2**x, the base-2 exponential of x.
	fmt.Println("math.Exp2(10)    =", math.Exp2(10))
	fmt.Println("math.Exp2(math.E)=", math.Exp2(math.E))

	/*
		Log returns the natural logarithm of x.
			Log(+Inf) = +Inf
			Log(0) = -Inf
			Log(x < 0) = NaN
			Log(NaN) = NaN
	*/
	x1 := math.Log(1)
	fmt.Printf("%.1f\n", x1)

	y1 := math.Log(2.7183)
	fmt.Printf("%.1f\n", y1)

	// Log10 returns the decimal logarithm of x.
	// The special cases are the same as for Log.
	fmt.Printf("%.1f\n", math.Log10(100))

	/*
		Sqrt returns the square root of x.
		func Sqrt(x float64) float64
			Sqrt(+Inf) = +Inf
			Sqrt(±0) = ±0
			Sqrt(x < 0) = NaN
			Sqrt(NaN) = NaN
	*/
	fmt.Println("math.Sqrt(4) =", math.Sqrt(4))
	fmt.Println("math.Sqrt(5) =", math.Sqrt(5))
	fmt.Println()

	fmt.Println("math.Cbrt(4) =", math.Cbrt(4))
	fmt.Println("math.Cbrt(5) =", math.Cbrt(5))
}
