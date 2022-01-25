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
	fmt.Printf("Hello! Mr.%s with age %d !!\n", name, age)
}

func main() {
	// Function calls
	hello()

	// 1) Function calls by positional arguments

	// hello2() // not enough arguments in call to hello2
	hello2("Girish")

	// hello2("Girish", "Harish")  // too many arguments in call to hello2

	// 2) keyword arguments not possible
	// hello2(name ="Girish") // syntax error: unexpected =, expecting comma or )
	// hello2(name : "Girish") // syntax error: unexpected =, expecting comma or )

	// hello3("Girish", "Harish")
	hello3("Girish", 78)

}
