package main

import "fmt"

func main() {
	// pointers are references to defined variables
	// some operations like 'call by reference' can be done based on pointer

	// every variable is a memory location and every memory location has its address

	var num1 int = 123 /* actual variable declaration */

	fmt.Printf("\nvalue            : %v", num1)
	fmt.Printf("\ntype             : %T", num1)
	fmt.Printf("\naddress location : %x", &num1) // c00000a0b0

	// pointer is a variable whose value is the address of another variable,
	// i.e., direct address of the memory location.
	fmt.Println()

	var num1_ptr *int                               /* pointer variable declaration */
	fmt.Printf("\nvalue            : %v", num1_ptr) // <nil>
	fmt.Printf("\ntype             : %T", num1_ptr)
	fmt.Printf("\naddress location : %x", num1_ptr)
	fmt.Println()

	num1_ptr = &num1                                /* store address of a in pointer variable*/
	fmt.Printf("\nvalue            : %v", num1_ptr) // 0xc00000a0b0
	fmt.Printf("\ntype             : %T", num1_ptr)
	fmt.Printf("\naddress location : %x", num1_ptr)

	// How to access the value referenced by the pointer
	fmt.Println("\n*num1_ptr    =", *num1_ptr)

	fmt.Println("num1 == *num1_ptr :", num1 == *num1_ptr)
	fmt.Println("&num1 == num1_ptr :", &num1 == num1_ptr)
}
