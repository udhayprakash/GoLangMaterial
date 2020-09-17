package main

import "fmt"

// Package-level declaration
var c, python, java int

func main() {
	// Function-level declaration
	var i float64
	var java float64 = 2.3
	fmt.Println("c      =", c)
	fmt.Println("python =", python)
	fmt.Println("java   =", java)
	fmt.Println("i      =", i)
}

// NOTE: Function-level declarations will be prioritized, compared to package-level declarations
