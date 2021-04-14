package main

import "fmt"

// Package-level declaration
var name = "package-level"

func main() {
	// Function-level declaration
	fmt.Println("In main()  : name   =", name)

	var name = "function-level"
	fmt.Println("In main()  : name   =", name)

	myFunc()
}

func myFunc() {
	fmt.Println("In myFunc(): name   =", name)
}

// NOTE: Function-level declarations will be prioritized,
// compared to package-level declarations
