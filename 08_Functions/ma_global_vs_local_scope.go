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

func incr(val int) {
	val++
}
func incrPtr(valPtr *int) {
	*valPtr++
}

func main() {
	v := 10
	incr(v)
	fmt.Println("After incr(v), v=", v) // 10

	incrPtr(&v)
	fmt.Println("After incrPtr(&v), v=", v) // 11

	// xpt := &v  // to createa pointer to a variable

	xPtr := new(int)
	// new() - takes type as argument, allocates enough memory
	//  to fit a value of that type, and returns a pointer to it
	fmt.Println("xPtr == nil :", xPtr == nil) // false
	fmt.Println("xPtr = ", xPtr)              // 0xc000014098
	fmt.Println("*xPtr = ", *xPtr)            // 0

	incrPtr(xPtr)
	fmt.Println("After incrPtr(xPtr), *xPtr=", *xPtr) // 1

}
