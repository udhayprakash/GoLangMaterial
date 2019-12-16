package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Changing value after declaration

	var x string
	x = "first"
	fmt.Println("x = ", x)

	x = "second"
	fmt.Println("x = ", x)

	x = x + "third"
	fmt.Println("x = ", x)

	x += "fourth" // Compound Operation
	fmt.Println("x = ", x)

	y := 5
	y += 1 // ; - statement separator
	fmt.Println("y = ", y)

	fmt.Println("reflect.TypeOf(x) =", reflect.TypeOf(x))
	fmt.Println("reflect.TypeOf(y) =", reflect.TypeOf(y))
}
