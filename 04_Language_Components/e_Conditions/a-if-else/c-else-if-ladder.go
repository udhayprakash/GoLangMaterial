package main

import (
	"fmt"
)

// Display largest number, among two numbers

func main() {
	var num1, num2 int
	fmt.Println("Please enter two numbers, space separated:")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	if num1 == num2 {
		fmt.Println("Both are equal")
	} else if num1 > num2 {
		fmt.Println("num1 is greater")
	} else { // num1 < num2
		fmt.Println("num2 is greater")
	}

	// var num3 int = 4 
	// var num4 float32 = 4.0
	// fmt.Println("num3 == num4:", num3 == num4)
	// incompatible types in binary expression
}

// Assignment - Display largest number, among three numbers
 