package main

import "fmt"

/*
Purpose: Scope - Global and local

	call by value
		- changes within the function will NOT reflect at the global level
		- arrays, basic types

	call by reference
		- changes within the function will reflect at the global level
		- slice, pointers, map, function or channel
*/

var pi float32 = 3.141

const PI float32 = 3.141

func case1() {
	fmt.Println("pi               = ", pi)
	fmt.Println("pi * 12          = ", pi*12)

	fmt.Println("PI               = ", PI)
	fmt.Println("PI * 12          = ", PI*12)
}

func case2() {
	// local changes will be reflected globally
	fmt.Println("\n\nbefore change pi = ", pi)
	pi = 3333333
	fmt.Println("after  change pi = ", pi)

	//fmt.Println("before change PI = ", PI)
	//PI = 3333333 // cant change value of const variable
	//fmt.Println("after  change PI = ", PI)
}

func case3(pi float32) {
	// Local changes wont affect global variables
	fmt.Println("\n\nbefore change pi = ", pi)
	pi = 444444
	fmt.Println("after  change pi = ", pi)

}
func main() {
	case1()
	case2()
	fmt.Println("Outside change pi = ", pi)
	//fmt.Println("Outside change PI = ", PI)

	case3(pi)
	fmt.Println("Outside change pi = ", pi)

}

// NOTE:
//1. Global variables can be used within functions, without passing as arguments
