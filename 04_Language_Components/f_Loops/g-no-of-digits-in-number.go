package main

import "fmt"

/*
Get number of digits in number
	input : 8768 --> output: 4
	input : 876  --> output: 3
*/

func main() {
	var num int
	fmt.Print("Enter an integer value:")

	success, err := fmt.Scanf("%d", &num)
	fmt.Println("success= ", success) // 1 if scanned, else 0
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		// Count the digits, in the number

		noOfDigits, sumOfDigits := 0, 0
		for num != 0 {
			sumOfDigits += num % 10 // lastDigit

			num = num / 10
			noOfDigits++
		}
		fmt.Println("noOfDigits =", noOfDigits)
		fmt.Println("sumOfDigits =", sumOfDigits)
	}

}
