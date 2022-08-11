package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Ceil returns the least integer value greater than or equal to x.
			Ceil(±0) = ±0
			Ceil(±Inf) = ±Inf
			Ceil(NaN) = NaN
	*/
	fmt.Printf("math.Ceil(1.01)         = %.2f\n", math.Ceil(1.01)) // 2.00
	fmt.Printf("math.Ceil(1.49)         = %.2f\n", math.Ceil(1.49)) // 2.00
	fmt.Printf("math.Ceil(1.51)         = %.2f\n", math.Ceil(1.51)) // 2.00
	fmt.Printf("math.Ceil(1.99)         = %.2f\n", math.Ceil(1.99)) // 2.00
	fmt.Println()

	/*
		Floor returns the greatest integer value less than or equal to x.
			Floor(±0) = ±0
			Floor(±Inf) = ±Inf
			Floor(NaN) = NaN
	*/
	fmt.Printf("math.Floor(1.01)        = %.2f\n", math.Floor(1.01)) // 1.00
	fmt.Printf("math.Floor(1.49)        = %.2f\n", math.Floor(1.49)) // 1.00
	fmt.Printf("math.Floor(1.51)        = %.2f\n", math.Floor(1.51)) // 1.00
	fmt.Printf("math.Floor(1.99)        = %.2f\n", math.Floor(1.99)) // 1.00
	fmt.Println()

	// Truncate will truncate the decimal value
	fmt.Printf("math.Trunc(1.01)        = %.2f\n", math.Trunc(1.01)) // 1.00
	fmt.Printf("math.Trunc(1.49)        = %.2f\n", math.Trunc(1.49)) // 1.00
	fmt.Printf("math.Trunc(1.51)        = %.2f\n", math.Trunc(1.51)) // 1.00
	fmt.Printf("math.Trunc(1.99)        = %.2f\n", math.Trunc(1.99)) // 1.00
	fmt.Println()

	/*
		Round returns the nearest integer, rounding half away
		from zero.
			Round(±0) = ±0
			Round(±Inf) = ±Inf
			Round(NaN) = NaN
	*/
	fmt.Printf("math.Round(1.01)        = %.2f\n", math.Round(1.01))  // 1.00
	fmt.Printf("math.Round(1.49)        = %.2f\n", math.Round(1.49))  // 1.00
	fmt.Printf("math.Round(1.51)        = %.2f\n", math.Round(1.51))  // 2.00
	fmt.Printf("math.Round(1.99)        = %.2f\n", math.Round(1.99))  // 2.00
	fmt.Printf("math.Round(-1.99)       = %.2f\n", math.Round(-1.99)) // -2.00
	fmt.Println()

	/*
		RoundToEven returns the nearest integer, rounding ties to even.
			RoundToEven(±0) = ±0
			RoundToEven(±Inf) = ±Inf
			RoundToEven(NaN) = NaN
	*/
	fmt.Println("math.RoundToEven(12.5)  =", math.RoundToEven(12.5))  // 10 12 14 => 12
	fmt.Println("math.RoundToEven(11.5)  =", math.RoundToEven(11.5))  // 10 12 14 => 12
	fmt.Println("math.RoundToEven(10.5)  =", math.RoundToEven(10.5))  // 10 12 14 => 10
	fmt.Println("math.RoundToEven(-11.5) =", math.RoundToEven(-11.5)) // -14 -12 -10 => -12

}
