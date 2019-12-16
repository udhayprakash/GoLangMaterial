package main

import "fmt"

func main() {
	// Logical AND Operation
	fmt.Println("true && true 	=", true && true)
	fmt.Println("true && false 	=", true && false)
	fmt.Println("false && true 	=", false && true)
	fmt.Println("false && false =", false && false)
	fmt.Println()

	// Logical OR Operation
	fmt.Println("true || true 	=", true || true)
	fmt.Println("true || false 	=", true || false)
	fmt.Println("false || true 	=", false || true)
	fmt.Println("false || false =", false || false)
	fmt.Println()

	// Logical Negation
	fmt.Println("!true			=", !true)
	fmt.Println("!false			=", !false)
	fmt.Println()

	result := (true && false) || (false && true) || !(false && false)
	fmt.Println("result			=", result)

}