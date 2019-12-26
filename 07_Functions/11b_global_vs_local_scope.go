package main

import "fmt"

/*
Purpose: Scope - Global and local

	call by value
		- changes within the function will NOT reflect at the global level
		- when values are passed as func arguments

	call by reference
		- changes within the function will reflect at the global level
		- when pointers are passed as func arguments
*/


func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	// new() - takes type as argument, allocates enough memory
	//  to fit a value of that type, and returns a pointer to it

	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
