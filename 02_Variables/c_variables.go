package main

import "fmt"

var num1, Num2, num3 int = 1, 2, 3

func main() {
	// Local Scope variable
	num1 := "one"

	// Multiple type values inference
	var lang, version, isLatest = "Go", 1.15, true

	fmt.Println("num1    :", num1)
	fmt.Println("Num2    :", Num2)
	fmt.Println("lang    :", lang)
	fmt.Println("version :", version)
	fmt.Println("isLatest:", isLatest)
}

/*
NOTE: 
 1. unused global variables will be ignored.
	No compilation Error.
 2. Variables are of two types 
	a. Public(exportable) variables
		- all variables starting with capital letter
		- all unicode(non-english) variables
	b. Private variables
		- variables starting with small case letter
*/
