package main

import "fmt"

var num1, num2 int = 1, 2

func main() {
	// Local Scope variable
	num1 := "one"

	// Multiple type values inference
	var lang, version, isLatest = "Go", 1.15, true

	fmt.Println("num1    :", num1)
	fmt.Println("lang    :", lang)
	fmt.Println("version :", version)
	fmt.Println("isLatest:", isLatest)
}

// NOTE: unused global variables will be ignored. 
//       No compilation Error.
