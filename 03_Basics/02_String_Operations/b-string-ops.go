package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello", "world")

	// String Concatenation
	fmt.Println("Hello" + "world")
	fmt.Println("Hello" + "" + "world")
	fmt.Println("Hello" + " " + "world")
	fmt.Println("Hello" + " " + `world`)

	// fmt.Println("Hello" + ' ' + `world`)
	// (mismatched types untyped string and untyped rune)

	// string repetion
	// fmt.Println("Hello " * 3)
	// invalid operation: "Hello " * 3 (mismatched types untyped string and untyped int)

	fmt.Println('H', 'H'*3) // 72, 216
}
