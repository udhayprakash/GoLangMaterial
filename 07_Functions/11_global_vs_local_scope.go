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
	fmt.Println("before change pi = ", pi)
	pi = 3333333
	fmt.Println("after  change pi = ", pi)

	//fmt.Println("before change PI = ", PI)
	//PI = 3333333 // cant change value of const variable
	//fmt.Println("after  change PI = ", PI)
}

func main() {
	case1()
	case2()
	fmt.Println("Outside change pi = ", pi)
	fmt.Println("Outside change PI = ", PI)

}
