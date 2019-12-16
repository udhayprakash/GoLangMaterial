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
	z := "Hello world"
	fmt.Println("z = ", z)

}

