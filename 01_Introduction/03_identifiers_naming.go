package main

import "fmt"

func main() {

	x := "Udhay"
	fmt.Println("Student's name is ", x)

	// code will be better readable if meaningful names are given for variables
	name := "Udhay"
	fmt.Println("Student's name is ", name)

	studentName := "Udhay"
	fmt.Println("Student's name is ", studentName)

	studentName_2 := "Udhay"
	fmt.Println("Student's name is ", studentName_2)

	// 2ndName = "prakash" // can't start with number
}

// NOTE:
// 1. camelCasing is preferred for the identifiers
// 2. Only variables, functions or types whose names begin with a capital
//    letter are considered as exported: accessible from packages outside
//    the current package.
