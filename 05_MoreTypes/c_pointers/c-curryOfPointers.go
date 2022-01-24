package main

import "fmt"

func main() {
	var intVar int
	var pointerVar *int
	var pointerToPointerVar **int

	fmt.Println("\nINITIAL DEFAULTS:")
	fmt.Println("intVar             :", intVar)              // 0
	fmt.Println("pointerVar         :", pointerVar)          // nil
	fmt.Println("pointerToPointerVar:", pointerToPointerVar) // nil

	intVar = 100
	pointerVar = &intVar
	// pointerToPointerVar = &intVar // cannot use &intVar (type *int) as type **int in assignment
	pointerToPointerVar = &pointerVar

	fmt.Println("\nAFTER ASSIGNMENT:")
	fmt.Println("intVar             :", intVar)
	fmt.Println("pointerVar         :", pointerVar)
	fmt.Println("pointerToPointerVar:", pointerToPointerVar)

	fmt.Println()
	fmt.Println("&intVar             :", &intVar)
	fmt.Println("&pointerVar         :", &pointerVar)
	fmt.Println("&pointerToPointerVar:", &pointerToPointerVar)

	fmt.Println()
	fmt.Println("*pointerVar          :", *pointerVar)
	fmt.Println("*pointerToPointerVar :", *pointerToPointerVar)
	fmt.Println("**pointerToPointerVar:", **pointerToPointerVar)
}
