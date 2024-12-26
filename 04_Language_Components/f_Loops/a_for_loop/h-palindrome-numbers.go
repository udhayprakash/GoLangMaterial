package main

import (
	"fmt"
)

// Get palindrome number between 100 & 999
// Ex: 101, 111, ... 888, ...

func main() {

	for num := 100; num < 1000; num++ {

		// num := 6776
		// orgNum := num
		// revNum := 0

		// num := 6776

		num := num  // safe copy of the loop scope variable

		orgNum, revNum := num, 0

		for num != 0 {
			lastDigit := num % 10
			// fmt.Println("lastDigit is ", lastDigit)
			// fmt.Println("remaining", num/10)

			// revNum += lastDigit
			revNum = ((revNum * 10) + lastDigit)

			num /= 10 // num = num /10 compound operator
		}

		// fmt.Println("\norgNum =", orgNum)
		// fmt.Println("revNum =", revNum)

		if orgNum == revNum {
			fmt.Printf("%d is palindrome\n", orgNum)
		}
	}
}

/*
6776	6
677		67
		(6 * 10) + 7
67		677
		((6 * 10) + 7) * 10 + 7
6		6776
		(((6 * 10) + 7) * 10 + 7) * 10 + 6

*/

// NOTE: variable defined inside a loop has loop scope
// NOTE:  you're modifying the num variable inside the loop, but then using it again in the outer loop
