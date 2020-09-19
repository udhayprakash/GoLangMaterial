package main

import "fmt"

func main() {
	// Changing value after declaration
	var x string
	x = "first"
	fmt.Println("x=", x)

	x = "second"

	// NOTE: Modified value should be of same data type

	//x = 44 // cannot use 44 (type untyped int) as type string in assignment
	//x := 44  // no new variables on left side of :=

	//var x float32 = 4.4 // x redeclared in this block
	//NOTE: A variable cant be declared more than once

}
