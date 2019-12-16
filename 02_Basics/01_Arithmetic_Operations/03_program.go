package main

import "fmt"

func main(){
	// Understanding the type- casting in division operation

	fmt.Println("10 / 3 	= ", 10 / 3) // int/int = int
	fmt.Println("10 / 3.0 	= ", 10 / 3.0) // int/float = float

	fmt.Println("10 % 3 	= ", 10 % 3) // Remainder
	// fmt.Println("10 % 3.0 = ", 10 % 3.0) // invalid operation: operator % not defined on untyped float
}