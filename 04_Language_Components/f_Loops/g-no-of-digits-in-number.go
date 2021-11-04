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
	}

    // Count the digits, in the number
	fmt.Printf("\nNumber %d ", num)
    sumOfDigits := 0
    noOfDigits := 0
    for num != 0 {
        // fmt.Println("\nnum /10=", num /10)
		// fmt.Println("num %10=", num %10)
        sumOfDigits = sumOfDigits + (num %10)

        num = num /10
        noOfDigits++
    }
    fmt.Printf("has %d digits\n", noOfDigits)
    fmt.Println("Sum of digits =", sumOfDigits)

}
 