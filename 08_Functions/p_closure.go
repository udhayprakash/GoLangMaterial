package main

import "fmt"

/*
Purpose: Closures
	-  functions which enclose other functions, often with
		a function return type.
	- The inner functions are able to reference the variables
		around the closure function.
NOTE: Inner functions are most often anonymous functions (or function literals).
*/

func callingHello() func() string {
	fmt.Println("starting callingHello ...")
	return func() string { // anonymous function
		fmt.Println("starting inner function ...")
		return "Hello world"
	}
}

func main() {
	fmt.Println("starting main ...")
	fmt.Println("callingHello =", callingHello)
	callingHello()

	result := callingHello()

	fmt.Println("\ncalling result ...")
	fmt.Println("result() = ", result())
}
