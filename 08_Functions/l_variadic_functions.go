package main

import "fmt"

/*
variadic functions
	- functions which support any no. of arguments
	- spread operator
*/

func oneArgFunc(num int) {
	fmt.Println("Args:", num)
}

func AnyNoOfArgs(num ...int) {
	fmt.Println("Args:", num)
}

func AnyNoOfArgsAnyType(num ...interface{}) {
	fmt.Println("Args:", num)
}

func main() {
	oneArgFunc(111) // Args: 111

	AnyNoOfArgs()
	AnyNoOfArgs(1)
	AnyNoOfArgs(1, 2)
	AnyNoOfArgs(1, 2, 3, 5, 6, 3, 2, 213)

	AnyNoOfArgsAnyType(1, '2', "udhay", true, [3]int{1, 2, 3})
	fmt.Println()

	Sum(12, 34, 67)
	fmt.Println()

	getFullName("udhay")
	getFullName("udhay", "prakash")
	getFullName("udhay", "prakash", "pethakamsetty")

}

func Sum(nums ...int) {
	result := 0
	for _, num := range nums {
		result += num
	}
	fmt.Printf("Sum(%d) = %d\n", nums, result)
}

func getFullName(name_parts ...string) {
	fullName := ""
	for _, name := range name_parts {
		fullName += name
	}
	fmt.Printf("getFullName(%v)=%v\n", name_parts, fullName)
}

// Assignment: Try to compute mean, median and mode
