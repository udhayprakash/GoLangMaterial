package main

import (
	"fmt"
	"math"
)

func main() {
	// Constants
	fmt.Println("Constants ...")
	fmt.Println("math.E       = ", math.E)
	fmt.Println("math.Pi      = ", math.Pi)
	fmt.Println("math.Phi     = ", math.Phi)
	fmt.Println("math.Sqrt2   = ", math.Sqrt2)
	fmt.Println("math.SqrtE   = ", math.SqrtE)
	fmt.Println("math.SqrtPi  = ", math.SqrtPi)
	fmt.Println("math.SqrtPhi = ", math.SqrtPhi)
	fmt.Println("math.Ln2     = ", math.Ln2)
	fmt.Println("math.Ln10    = ", math.Ln10)
	fmt.Println("math.Log2E   = ", math.Log2E)
	fmt.Println("math.Ln10    = ", math.Ln10)
	fmt.Println("math.Log10E  = ", math.Log10E)

	fmt.Println("\n\nFloating -point value limits ...")
	fmt.Println("math.MaxFloat32             = ", math.MaxFloat32)
	fmt.Println("math.SmallestNonzeroFloat32 = ", math.SmallestNonzeroFloat32)
	fmt.Println("math.MaxFloat64             = ", math.MaxFloat64)
	fmt.Println("math.SmallestNonzeroFloat64 = ", math.SmallestNonzeroFloat64)

	fmt.Println("\n\nInteger value limits ...")
	fmt.Println("math.MaxInt8  = ", math.MaxInt8)
	fmt.Println("math.MinInt8  = ", math.MinInt8)
	fmt.Println("math.MaxInt16 = ", math.MaxInt16)
	fmt.Println("math.MinInt16 = ", math.MinInt16)
	fmt.Println("math.MaxInt32 = ", math.MaxInt32)
	fmt.Println("math.MinInt32 = ", math.MinInt32)
	fmt.Println("math.MaxInt64 = ", math.MaxInt64)
	fmt.Println("math.MinInt64 = ", math.MinInt64)

	fmt.Println("\nmath.MaxUint8 = ", math.MaxUint8)
	fmt.Println("math.MaxUint16= ", math.MaxUint16)
	fmt.Println("math.MaxUint32= ", math.MaxUint32)

	//fmt.Println("math.MaxUint64= ", math.MaxUint64) //PANIC: constant 18446744073709551615 overflows int

	var value uint64 = math.MaxUint64
	fmt.Println("math.MaxUint64= ", value)
}
