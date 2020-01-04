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

	studentName2 := "Udhay"
	fmt.Println("Student's name is ", studentName2)

	// 2ndName = "prakash" // can't start with number

	// To get the address location where an object is stored
	fmt.Println(studentName, &studentName)
	fmt.Println(studentName2, &studentName2)

	fmt.Println()
	emptyString1  := ""
	var emptyString2 string
	var emptyString3 = ""
	var emptyString4 string = ""
	fmt.Println("emptyString1 == emptyString2 :", emptyString1 == emptyString2)
	fmt.Println("emptyString1 == emptyString3 :", emptyString1 == emptyString3)
	fmt.Println("emptyString1 == emptyString4 :", emptyString1 == emptyString4)
}

// NOTE:
// 1. camelCasing is preferred for the identifiers
// 2. Only variables, functions or types whose names begin with a capital
//    letter are considered as exported: accessible from packages outside
//    the current package.
// 3. Names starting with a capital letter are public, names with a
//    lowercase letter are private.

