/*
Go natively handles Unicode, so
it can process text in all the world’s languages.
*/
package main

import "fmt"

func main() {
	fmt.Println('H') // 72  - because 72 is ascii code for 'H'
	//fmt.Println('Hello world') // more than one character in rune literal
	// single quotes can't be used

	// double Quote Strings - - single line strings
	fmt.Println("H") // H
	fmt.Println("Hello go world")
	fmt.Println("Hello\tgo\nworld")
	fmt.Println("Hello\bgo\nworld")
	fmt.Println("123\bDOG")

	fmt.Println("Quotes: \"\"")
	fmt.Println("Backslash: \\")

	//fmt.Println("Hello
	//			 go world")

	// back-tick (raw) strings - Both Single & Multi-line Strings
	fmt.Println(`H`) // H
	fmt.Println(`Hello go world`)
	fmt.Println(`Hello\tgo\nworld`) // Escape sequences were not recognized
	fmt.Println(`Quotes: ""`)
	fmt.Println(`Backslash: \\`)
	fmt.Println(`Hello 
                      world`)
}
