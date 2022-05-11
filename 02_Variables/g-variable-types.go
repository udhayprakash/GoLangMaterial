package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string

	// changing value after declaration
	x = "first"          // updating
	fmt.Println("x=", x) // first

	x = "second"         // updating
	fmt.Println("x=", x) // second

	x = x + "third"      // strng concatenation
	fmt.Println("x=", x) // secondthird

	x += "fourth"        // strng concatenation
	fmt.Println("x=", x) // secondthirdfourth

	fmt.Println()

	fmt.Println("x                   	   =", x)
	fmt.Println("reflect.ValueOf(x)  	   =", reflect.ValueOf(x))
	fmt.Println("reflect.TypeOf(x)   	   =", reflect.TypeOf(x))

	fmt.Println("reflect.TypeOf(x).String()=", reflect.TypeOf(x).String())
	fmt.Println("reflect.TypeOf(x).Kind()  =", reflect.TypeOf(x).Kind())
	// NOTE: TypeOf and Kind will be same for basic data types.
	//       Will differ for structs, ...

	fmt.Println()

	// y := 5
	// y += 1

	y := 5
	y += 1 // ; - statement separator
	fmt.Println("y =", y)
	fmt.Println("reflect.TypeOf(y)   	=", reflect.TypeOf(y))
	fmt.Println()

	fmt.Println("reflect.TypeOf(42)  	=", reflect.TypeOf(42))
	fmt.Println("reflect.TypeOf(42.0)	=", reflect.TypeOf(42.0))
	fmt.Println("reflect.TypeOf(4223434234343242342) =", reflect.TypeOf(4223434234343242342))
	fmt.Println("reflect.TypeOf('a') 	=", reflect.TypeOf('a'))  // int32
	fmt.Println("reflect.TypeOf(\"a\")	=", reflect.TypeOf("a")) // string
	fmt.Println("reflect.TypeOf(true)	=", reflect.TypeOf(true)) // bool
	fmt.Println("reflect.TypeOf(nil)	=", reflect.TypeOf(nil))   //  <nil>
	fmt.Println()

	z := 'a'
	fmt.Println("reflect.TypeOf(z).String()=", reflect.TypeOf(z).String())
	fmt.Println("reflect.TypeOf(z).Kind()  =", reflect.TypeOf(z).Kind())
	fmt.Println()

	fmt.Printf("%T\n", 0)
	fmt.Printf("value is %v and type is %T\n", 0, 0)
	fmt.Println()

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

	// decimal, binary & Hexadecimal representation
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

}
