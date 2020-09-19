package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Constants may be typed or untyped
	var x int = 10
	fmt.Println("x 		=", x)

	x = 20
	fmt.Println("x 		=", x)

	// `const` declares a constant value.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const y int = 10
	fmt.Println("y		=", y)

	//y = 20 // cannot assign to y
	// NOTE: cannot update value to constant variables

	// Constant expressions perform arithmetic with arbitrary precision.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(n, reflect.TypeOf(n))
	fmt.Println(d, reflect.TypeOf(d))

	// Explicit type conversion.
	fmt.Println(int64(d), reflect.TypeOf(int64(d)))
	fmt.Println()

	// An untyped constant has no limits. When it's used in a context
	// that requires a type, a type will be inferred and a limit applied.
	// Too big for an int
	const big = 10000000000
	fmt.Println("big    =", big, reflect.TypeOf(big))

	// Still ok: "untyped = untyped * untyped"
	const bigger = big * 100
	fmt.Println("bigger =", bigger, reflect.TypeOf(bigger))

	// No problem: Result fits in an int
	var i int = big / 100
	fmt.Println("i 		=", i, reflect.TypeOf(i))

	// Compile time error: "constant 10000000000 overflows int"
	//var j int16 = big
	//fmt.Println("j =", j, reflect.TypeOf(j))
}
