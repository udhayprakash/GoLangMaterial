package main

import (
	"fmt"
	"math"
)

func main(){
	/*
		Ceil returns the least integer value greater than or equal to x.
			Ceil(±0) = ±0
			Ceil(±Inf) = ±Inf
			Ceil(NaN) = NaN
	*/
	c := math.Ceil(1.49)
	fmt.Printf("%.1f\n", c)

	/*
		Floor returns the greatest integer value less than or equal to x.
			Floor(±0) = ±0
			Floor(±Inf) = ±Inf
			Floor(NaN) = NaN
	*/
	c1 := math.Floor(1.51)
	fmt.Printf("%.1f\n", c1)


	/*
		Round returns the nearest integer, rounding half away
		from zero.
			Round(±0) = ±0
			Round(±Inf) = ±Inf
			Round(NaN) = NaN
	*/
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p) // 11

	p = math.Round(10.4)
	fmt.Printf("%.1f\n", p) // 10.0

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)


	/*
		RoundToEven returns the nearest integer, rounding ties to even.
			RoundToEven(±0) = ±0
			RoundToEven(±Inf) = ±Inf
			RoundToEven(NaN) = NaN
	*/
	fmt.Println("math.RoundToEven(12.5) =", math.RoundToEven(12.5)) // 10 12 14 => 12
	fmt.Println("math.RoundToEven(11.5) =", math.RoundToEven(11.5)) // 10 12 14 => 12
	fmt.Println("math.RoundToEven(10.5) =", math.RoundToEven(10.5)) // 10 12 14 => 10
	fmt.Println("math.RoundToEven(-11.5)=", math.RoundToEven(-11.5))// -14 -12 -10 => -12

}
