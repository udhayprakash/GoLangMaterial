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


// Function Definition
func hello(name string){
	fmt.Println("Hello" , name)
}

func main()  {
	// Function call
	//hello() // not enough arguments in call to hello
	hello("Python") //Hello Python

	//hello("python", "program") // too many arguments in call to hello

	// keyword arguments not possible
	//hello(name="python")  // syntax error


}