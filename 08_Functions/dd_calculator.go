package main

import "fmt"

// Function Definitions
func add(a float32, b float32) float32 {
	return a + b
}

func subtract(a, b float32) float32 {
	return a - b
}

func multiply(a, b float32) float32 {
	return a * b
}

func divide(a, b float32) float32 {
	return a / b
}

func main() {
	// Declare variables
	var num1, num2 float32
	var operator string

	fmt.Print("Enter No 1: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter No 2: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println("Your answer is: ", add(num1, num2))
	case "-":
		fmt.Println("Your answer is: ", subtract(num1, num2))
	case "*":
		fmt.Println("Your answer is: ", multiply(num1, num2))
	case "/":
		fmt.Println("Your answer is: ", divide(num1, num2))
	default:
		fmt.Println("Invalid operator")
	}

}
