package main

import "fmt"

// Package-level declaration
var name = "package-level"

func main() {
	fmt.Println("In main()  : name   =", name) // package-level

	// Function-level declaration
	var name = "function-level"
	fmt.Println("In main()  : name   =", name) // function-level

	MyFunc()
}

func MyFunc() {
	fmt.Println("In myFunc(): name   =", name) // package-level
}

// NOTE: Function-level declarations will be prioritized,
// compared to package-level declarations
