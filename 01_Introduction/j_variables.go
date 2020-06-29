package main

import "fmt"

/*
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted;
the variable will take the type of the initializer.
*/
var num1, num2 int = 1, 2

func main(){
	// Local Scope variable
	num1 := "one"

	// Multiple type values inference
	var lang, version, isLatest = "Go", 1.14, true

	fmt.Println("num1    :", num1)
	fmt.Println("lang    :", lang)
	fmt.Println("version :", version)
	fmt.Println("isLatest:", isLatest)

}

// unused global variables will be ignored. No compilation Error.