package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 10
	fmt.Println("x=", x)

	x = 20
	fmt.Println("x=", x)

	// `const` declares a constant value.
	const y int = 10
	fmt.Println("y=", y)

	// y = 20				// cannot update value to constant variables
	// fmt.Println("y=", y)

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(n, reflect.TypeOf(n))
	fmt.Println(d, reflect.TypeOf(d))

	// Explicit type conversion.
	fmt.Println(int64(d), reflect.TypeOf(int64(d)))
}
