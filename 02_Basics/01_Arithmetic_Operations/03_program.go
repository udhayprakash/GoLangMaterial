package main

import "fmt"

func main() {
	// Understanding the type- casting in division operation

	fmt.Println("10 / 3  = ", 10/3)   // int/int = int
	fmt.Println("10 / 3.0= ", 10/3.0) // int/float = float

	fmt.Println("10 % 3  = ", 10%3) // Remainder
	// fmt.Println("10 % 3.0 = ", 10 % 3.0) // invalid operation: operator % not defined on untyped float

	// Another example
	fmt.Println("5.0/4.0 =", 5.0/4.0) // 1.25
	fmt.Println("5  /4   =", 5/4)     // 1
	// Because integer division truncates the result towards zero

	// sign of remainder is always same as sign of the dividend
	fmt.Println("-5/3   =", -5/3)  // -1
	fmt.Println("-5/-3  =", -5/-3) // 1

}
