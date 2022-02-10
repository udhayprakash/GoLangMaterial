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
	pi = 333
	fmt.Println("after  change pi = ", pi)

	//fmt.Println("before change PI = ", PI)
	//PI = 3333333 // cant change value of const variable
	//fmt.Println("after  change PI = ", PI)
}

func case3(pi float32) {
	// Local changes wont affect global variables
	fmt.Println("\n\nbefore change pi = ", pi)
	pi = 444
	fmt.Println("after  change pi = ", pi)

}
func case4Maps(m map[int64]int) {
	m[42] = 24
}

func case5Slice(m []int) {
	m = append(m, 42)
}

func case6Array(m [4]int) {
	m[2] = 333
	fmt.Println("In case6Array, m=", m) // [11 22 333 44]
}

func main() {
	case1()

	case2()
	fmt.Println("Outside change pi = ", pi)
	//fmt.Println("Outside change PI = ", PI)

	case3(pi)
	fmt.Println("Outside change pi = ", pi)

	// Maps - changes within function are reflected outside
	data := make(map[int64]int)
	case4Maps(data)
	fmt.Println(data) // map[42:24]

	// Slice - changes within function are NOT reflected outside
	mySlice := make([]int, 0, 20)
	case5Slice(mySlice)
	fmt.Println(mySlice) // []

	// Array - changes within function are NOT reflected outside
	myArray := [4]int{11, 22, 33, 44}
	case6Array(myArray)
	fmt.Println("\n myArray =", myArray) //  [11 22 33 44]
}

// NOTE:
//1. Global variables can be used within functions, without passing as arguments
