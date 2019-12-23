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
	fmt.Println()

	fmt.Println("x                   =", x)
	fmt.Println("reflect.ValueOf(x)  =", reflect.ValueOf(x))
	fmt.Println("reflect.TypeOf(x)   =", reflect.TypeOf(x))
	fmt.Println("reflect.TypeOf(x).Kind()=", reflect.TypeOf(x).Kind())
	// TypeOf and Kind will be same for basic data types. Will differ for structs, ...
	fmt.Println()

	fmt.Println("reflect.TypeOf(y)   =", reflect.TypeOf(y))

	fmt.Println("reflect.TypeOf(42)  =", reflect.TypeOf(42))
	fmt.Println("reflect.TypeOf(42.0)=", reflect.TypeOf(42.0))
	fmt.Println("reflect.TypeOf(4223434234343242342) =", reflect.TypeOf(4223434234343242342))
	fmt.Println("reflect.TypeOf('a') =", reflect.TypeOf('a'))   // int32
	fmt.Println("reflect.TypeOf(\"a\")=", reflect.TypeOf("a")) // string
	fmt.Println("reflect.TypeOf(true)=", reflect.TypeOf(true)) // bool

}
