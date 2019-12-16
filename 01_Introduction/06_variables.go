package main

import "fmt"

func main() {
	// Changing value after declaration

	var x string 
	x = "first"
	fmt.Println("x = ", x)

	// NOTE: Modified value should be of same data type
	// x = 4
	// fmt.Println("x = ", x)

	// x :=  4.4
	// fmt.Println("x = ", x)

	// var x bool = true         // x redeclared in this block
	// fmt.Println("x = ", x)
}

