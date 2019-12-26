package main

import "fmt"
// Ensure that the type of the variable can handle all the values
// of business logic
func main() {
	var blue int = 255
	fmt.Printf("blue: %3d BinaryForm: %b\n", blue, blue)
	blue += 1
	fmt.Printf("blue: %3d BinaryForm: %b\n", blue, blue)
	fmt.Println()

	//var green int8 = 255  // constant 255 overflows int8
	var green uint8 = 255
	fmt.Printf("green: %3d BinaryForm: %b\n", green, green)
	green += 1
	fmt.Printf("green: %3d BinaryForm: %b\n", green, green)

	//  11111111
	// +       1
	// ----------
	//  00000000

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"
}
