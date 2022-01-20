package main

import "fmt"

func main() {
	var blue int = 255
	fmt.Println("blue= ", blue)
	fmt.Printf("%7T  blue: %d BinaryForm: %3b\n", blue, blue, blue)

	blue += 1 // blue = blue + 1
	fmt.Printf("%7T  blue: %d BinaryForm: %3b\n", blue, blue, blue)


	// var green int8 = 255 // error: integer constant overflow
	// constant 255 overflows int8  -> -127 to 128  --> 2^8 - 1 = 255

	
	var green uint8 = 255   // uint8: 0 to 255
	fmt.Printf("%7T green: %3d BinaryForm: %3b\n", green, green, green)

	green += 1
	fmt.Printf("%7T green: %3d BinaryForm: %b\n", green, green, green)

	
	//  11111111
	// +       1
	// ----------
	//  00000000

	
	var i int8 = 127
	fmt.Println(i -2, i-1, i, i+1, i+2, i+3) // 125 126 127 -128 -127 -126

	// uint8 : 0 to 255
	// int8  : -127 to 127
}
