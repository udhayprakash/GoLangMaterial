package main

import (
	"fmt"
	"math/bits"
)

// Main function
func main() {

	// Using Reverse() function
	a := bits.Reverse(5)
	fmt.Printf("Original value: %d \n Binary Form: %b\n", 5, 5)
	fmt.Printf("Reverse  value: %d \n Binary Form: %b\n", a, a)

	// subtract bits  diff = a – b – borrow
	nvalue_1, borrowOut := bits.Sub(4, 3, 0)
	fmt.Println("Diff:", nvalue_1)
	fmt.Println("BorrowOut :", borrowOut)
}
