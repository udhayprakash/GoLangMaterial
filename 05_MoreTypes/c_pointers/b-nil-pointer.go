package main

import (
	"fmt"
)

func main() {
	var ptr *float64                               // nil pointer
	fmt.Printf("The value of ptr is : %x\n", ptr)  // 0
	fmt.Printf("\ntype              : %T\n", ptr)  // *float64
	fmt.Println("Before ptr != nil :", ptr != nil) // false

	// var num2 float64 = 123.213
	num2 := 123.213

	// referencing
	ptr = &num2
	fmt.Println("After  ptr != nil :", ptr != nil) // true

	// dereferencing
	ptr = nil
	fmt.Println("After  ptr != nil :", ptr != nil) // false

	
	// myStr := "Golang"
	// ptr = &myStr
	// cannot use &myStr (type *string) as type *float64 in assignment

	// num3 := 123
	// ptr = &num3
	// cannot use &num3 (type *int) as type *float64 in assignment

	// var num4 float32 = 1.3
	// ptr = &num4
	// cannot use &num4 (type *float32) as type *float64 in assignment

	// var floatPtr *float  //  undefined: float
}

// NOTE: You can refer values of same data types as defined in pointer
