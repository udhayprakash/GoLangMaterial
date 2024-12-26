package main

import "fmt"

// Display largest number, among two given numbers

func main() {

	var num1, num2 int

	fmt.Println("Please enter two numbers, space separated:")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	fmt.Printf(`
			num1 = %d
			num2 = %d
	`, num1, num2)

	// Method 1
	if num1 > num2 {
		fmt.Println("num1 is greater")
	} else if num1 < num2 {
		fmt.Println("num2 is greater")
	} else { // num1 == num2
		fmt.Println("Both are equal")
	}

	// Method 2
	if num1 == num2 {
		fmt.Println("Both are equal")
	} else if num1 > num2 {
		fmt.Println("num1 is greater")
	} else { //  num1 < num2
		fmt.Println("num2 is greater")
	}

}

// Assignment - Display largest number, among three numbers
