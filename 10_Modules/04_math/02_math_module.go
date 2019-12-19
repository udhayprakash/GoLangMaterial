package main

import (
	"fmt"
	"math"
)

func main() {

	/*
	Abs(±Inf) = +Inf
	Abs(NaN) = NaN
	 */
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)

	// Acos(x) = NaN if x < -1 or x > 1
	fmt.Printf("%.2f", math.Acos(1))

	/*
	Acosh(+Inf) = +Inf
	Acosh(x) = NaN if x < 1
	Acosh(NaN) = NaN
	*/
	fmt.Printf("%.2f", math.Acosh(1))

	/*
	Asin(±0) = ±0
	Asin(x) = NaN if x < -1 or x > 1
	*/
	fmt.Printf("%.2f", math.Asin(0))

	/*
	Asinh(±0) = ±0
	Asinh(±Inf) = ±Inf
	Asinh(NaN) = NaN
	*/
	fmt.Printf("%.2f", math.Asinh(0))


	/*
	Atan(±0) = ±0
	Atan(±Inf) = ±Pi/2
	*/
	fmt.Printf("%.2f", math.Atan(0))


	/*
	Ceil returns the least integer value greater than or equal to x.
		Ceil(±0) = ±0
		Ceil(±Inf) = ±Inf
		Ceil(NaN) = NaN
	*/
	c := math.Ceil(1.49)
	fmt.Printf("%.1f\n", c)

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
	fmt.Println("math.Dim(7, 3) =", math.Dim(7, 3))
	fmt.Println("math.Dim(3, 7) =", math.Dim(3, 7))
	fmt.Println("math.Dim(3, 3) =", math.Dim(3, 3))

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
	Floor returns the greatest integer value less than or equal to x.
		Floor(±0) = ±0
		Floor(±Inf) = ±Inf
		Floor(NaN) = NaN
	 */
	c1 := math.Floor(1.51)
	fmt.Printf("%.1f", c1)

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
	Max returns the larger of x or y.
		Max(x, +Inf) = Max(+Inf, x) = +Inf
		Max(x, NaN) = Max(NaN, x) = NaN
		Max(+0, ±0) = Max(±0, +0) = +0
		Max(-0, -0) = -0
	*/
	fmt.Println("math.Max(7, 3)=", math.Max(7, 3))
	fmt.Println("math.Max(3, 7)=", math.Max(3, 7))
	fmt.Println("math.Max(3,-7)=", math.Max(3, -7))

	/*
	Min returns the smaller of x or y.
		Min(x, -Inf) = Min(-Inf, x) = -Inf
		Min(x, NaN) = Min(NaN, x) = NaN
		Min(-0, ±0) = Min(±0, -0) = -0
	*/
	fmt.Println("math.Min(7, 3)=", math.Min(7, 3))
	fmt.Println("math.Min(3, 7)=", math.Min(3, 7))
	fmt.Println("math.Min(3,-7)=", math.Min(3, -7))


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
	c2 := math.Mod(7, 4)
	fmt.Printf("%.1f\n", c2)

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
	Round returns the nearest integer, rounding half away
	from zero.
		Round(±0) = ±0
		Round(±Inf) = ±Inf
		Round(NaN) = NaN
	*/
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)

	/*
	RoundToEven returns the nearest integer, rounding ties to even.
		RoundToEven(±0) = ±0
		RoundToEven(±Inf) = ±Inf
		RoundToEven(NaN) = NaN
	*/
	fmt.Println("math.RoundToEven(12.5) =", math.RoundToEven(12.5))
	fmt.Println("math.RoundToEven(11.5) =", math.RoundToEven(11.5))
	fmt.Println("math.RoundToEven(-11.5)=", math.RoundToEven(-11.5))

	// Signbit reports whether x is negative or negative zero.
	math.Signbit(10)

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
}
