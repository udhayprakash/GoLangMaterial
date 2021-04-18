package main

import "fmt"

// Calculator Application
func main() {

	// num1 := 100
	// num2 := 200

	var num1, num2 int32
	fmt.Println("Enter the values of num1 and num2:")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	fmt.Println("num1 + num2 = ", num1+num2)
	fmt.Println("num1 - num2 = ", num1-num2)
	fmt.Println("num1 * num2 = ", num1*num2)
	fmt.Println("num1 / num2 = ", num1/num2)

}
