package main

import (
	"fmt"
	"reflect"
)

const num = 10

func main() {
	fmt.Println("num	=", num)

	var x int = 10
	fmt.Println("x 		=", x)

	x = 20
	fmt.Println("x 		=", x)

	y := 30
	fmt.Println("y 		=", y)

	y = 40
	fmt.Println("y 		=", y)
	fmt.Println()

	// `const` declares a constant value.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const z int = 10
	fmt.Println("z		=", z)

	// z = 20  // cannot assign to z (constant 10 of type int)

	// Constant expressions perform arithmetic with arbitrary precision.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(n, reflect.TypeOf(n))
	fmt.Println(d, reflect.TypeOf(d))

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

	// var j int16 = big
	// cannot use big (untyped int constant 10000000000) as int16 value in variable declaration (overflows)

	var j int = big
	fmt.Println("j =", j)
}
