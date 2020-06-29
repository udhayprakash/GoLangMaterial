package main

import "fmt"

func main() {
	// Changing value after declaration

	var x string
	x = "first"
	fmt.Println("x = ", x)

	// NOTE: Modified value should be of same data type

	//x = 4 // cannot use 4 (type untyped int) as type string in assignment
	//x :=  4.4 // cannot use 4.4 (type untyped float) as type string in assignment

	//var x float32 = 4.4  // x redeclared in this block
	//NOTE: A variable cant be declared more than once
}