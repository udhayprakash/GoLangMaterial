package main

import "fmt"

// Package-level declaration
var name = "package-level"

func main() {
	fmt.Println("In main()  : name   =", name)

	// Function-level declaration
	var name = "function-level"

	fmt.Println("In main()  : name   =", name)

	MyFunc()
}

func MyFunc() {
	fmt.Println("In myFunc(): name   =", name)
}

// NOTE: Function-level declarations will be prioritized,
// compared to package-level declarations
