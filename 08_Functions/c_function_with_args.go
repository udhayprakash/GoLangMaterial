package main

import "fmt"

/*
Purpose: Function with one input argument and no return value

	function calls
		- by positional arguments
			- variables must be passed explicitly
			- Functions don't have access to anything in calling function,
              unless it's passed in explicitly
		- by keyword arguments NOT possible in golang
*/

// Function Definitions
func hello() {
	fmt.Println("Hello World")
}

func hello2(name string) {
	fmt.Println("Hello", name)
}

func hello3(name string, age int) {
	fmt.Println("Hello", name, "! Your age is", age)
}

func main() {
	// Function calls
	hello()

	// hello2() // not enough arguments in call to hello2
	hello2("Danilo")

	// hello2("Danilo", "Jo")  // too many arguments in call to hello2

	// // keyword arguments not possible
	// hello2(name = "Danilo") // syntax error: unexpected =, expecting comma or )

	hello3("Danilo", 29)
}
 