package main

import (
	"fmt"
)

func main() {
	var ptr *float64
	fmt.Printf("The value of ptr is : %x\n", ptr)

	num2 := 123.213

	ptr = &num2
	fmt.Println("After  ptr != nil :", ptr != nil)

	// myStr := "Golang"
	// ptr = &myStr
	// // error: incompatible types in assignment

	// num3 := 123
	// ptr = &num3
	// incompatible types in assignment
}


// NOTE: You can refer values of same data types as defined in pointer 