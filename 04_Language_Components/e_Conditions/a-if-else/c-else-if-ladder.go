package main

import (
	"fmt"
)

// Display largest number, among two numbers
func main() {
	var num1, num2 int
	fmt.Println("Please enter two number, space separated:")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	if num1 > num2 {
		fmt.Println("num1 is greater")
	} else if num2 > num1 {
		fmt.Println("num2 is greater")
	} else {
		fmt.Println("num1 is equal to num2")
	}
}

// Assignment - Display largest number, among three numbers