package main

import "fmt"

func main() {
	// Variable declaration and initialization

	// Method 1
	var x string
	x = "Hello world"
	fmt.Println("x = ", x)

	// Method 2
	var y string = "Hello world"
	fmt.Println("y = ", y)

	// Method 3 - type is not necessary because the Go compiler
	// is able to infer the type based on the literal value you
	// assign the variable.
	z := "Hello world"  // Type inference - variable's type is inferred from the value on the right hand side.
	fmt.Println("z = ", z)

	i := 42           // int
	f := 3.142        // float64 - an IEEE-754 64-bit floating point number.
	g := 0.867 + 0.5i // complex128 - represented internally with two float64's.
	fmt.Println(i, f, g)

}
