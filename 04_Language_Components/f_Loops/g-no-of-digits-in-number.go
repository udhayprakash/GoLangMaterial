package main

import "fmt"

/*
Get number of digits in number
	input : 8768 --> output: 4
	input : 876  --> output: 3
*/

func main() {
	var num int
	fmt.Println("Enter an integer value:")

	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Error - ", err)
	}

	// Count the digits
	digitsCount := 0
	sumOfDigits := 0
	fmt.Printf("Number %d ", num)
	for num != 0 {
		fmt.Println("\nnum /10=", num /10)
		fmt.Println("num %10=", num %10)
		sumOfDigits = sumOfDigits + (num % 10)
		num = num / 10
		digitsCount++
	}
	fmt.Printf("has %d digits\n", digitsCount)
	fmt.Println("Sum of digits =", sumOfDigits)

}
