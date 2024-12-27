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
	// pointerToPointerVar = &&intVar
	pointerToPointerVar = &pointerVar

	fmt.Println("\nAFTER ASSIGNMENT:")
	fmt.Println("intVar             :", intVar)              // 100
	fmt.Println("pointerVar         :", pointerVar)          // 0xc0000100a8
	fmt.Println("pointerToPointerVar:", pointerToPointerVar) // 0xc000006028

	fmt.Println()
	fmt.Println("&intVar             :", &intVar)
	fmt.Println("&pointerVar         :", &pointerVar)
	fmt.Println("&pointerToPointerVar:", &pointerToPointerVar)

	fmt.Println()
	fmt.Println("*pointerVar          :", *pointerVar)           // 100
	fmt.Println("*pointerToPointerVar :", *pointerToPointerVar)  // 0xc0000100a8
	fmt.Println("**pointerToPointerVar:", **pointerToPointerVar) // 100

	// Updating values
	*pointerVar = 200

	fmt.Println()
	fmt.Println("intVar               :", intVar)                // 200
	fmt.Println("*pointerVar          :", *pointerVar)           // 200
	fmt.Println("**pointerToPointerVar:", **pointerToPointerVar) // 200


	// Updating values
	**pointerToPointerVar = 300

	fmt.Println()
	fmt.Println("intVar               :", intVar)                // 300
	fmt.Println("*pointerVar          :", *pointerVar)           // 300
	fmt.Println("**pointerToPointerVar:", **pointerToPointerVar) // 300
}
