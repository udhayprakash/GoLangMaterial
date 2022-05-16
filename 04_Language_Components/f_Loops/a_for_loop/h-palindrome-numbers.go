package main

import (
	"fmt"
)

// Get palindrome number between 100 & 999
// Ex: 101, 111, ... 888, ... 999

func main() {
	// for num:= 0; num <= 999; num++{
	// }

	num := 6776
	orgNum := num
	revNum := 0
	for num != 0 {
		fmt.Println("\nnum /10=", num/10)
		fmt.Println("num %10=", num%10)

		revNum += (num % 10)

		num /= 10
	}

	fmt.Println("\norgNum =", orgNum)
	fmt.Println("revNum =", revNum)
}

//  6776
// 6000
//  700
//   70
//    6
