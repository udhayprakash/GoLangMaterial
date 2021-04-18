package main

import (
	"fmt"
)

/*
Go natively handles Unicode, so
it can process text in all the world’s languages.
*/

func main() {
	fmt.Println("H") // H
	fmt.Println('H') // 72  - because 72 is ascii code for 'H'
	fmt.Println("α") // α
	fmt.Println('α') // 945 -- UNICODE
	fmt.Println()

	fmt.Println("Hello World")
	// fmt.Println('Hello World')// more than one character in rune literal
	// single quotes can't be used

	// double Quote Strings -single line strings
	fmt.Println("Hello go world")
	fmt.Println("Hello\tgo\nworld")
	fmt.Println("Hello\bgo\nworld")
	fmt.Println("123\bDOG")
	fmt.Println("123456\rDOG")

	fmt.Println("Quotes: '")
	fmt.Println("Quotes: \"")
	fmt.Println("Backslash: \\")

	// Multi-line strings
	// fmt.Println("Hello
	// 			 world")

	// back-tick (raw) strings - Both Single & Multi-line Strings
	fmt.Println(`H`) // H
	fmt.Println(`Hello go world`)
	fmt.Println(`Hello\tgo\nworld`) // Escape sequences were not recognized
cls	fmt.Println(`Quotes: "`)
	fmt.Println(`Backslash: \\`)
	fmt.Println(`Hello 
                      world`)
}
