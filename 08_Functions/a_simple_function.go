package main

import (
	"fmt"
	"reflect"
)

/*
Purpose: Functions
            - code re-usability
            - To modularize the problem
            - Better maintenance of the code
			- treated as first class objects
			- function definitions need not be before calls

Go requires explicit returns, i.e. it won't automatically
return the value of the last expression.

SYNTAX:
	func name(parameter-list) (result-list) {
		body
	}

main()  and init() will be automatically called
*/

// Function defintion
func myFunc() {
	//Function with no arguments and no return value
	fmt.Println("Function myFunc is called!!!")
}

func main() {
	fmt.Println("myFunc                ", myFunc) // 0xc2b460
	fmt.Println("reflect.TypeOf(myFunc)", reflect.TypeOf(myFunc), reflect.TypeOf(myFunc).Kind())

	// function call
	myFunc()

	// fmt.Println("myFunc() = ", myFunc()) // Error as no return from function
}
