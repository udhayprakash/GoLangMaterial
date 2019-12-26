package main

import "fmt"

// Go natively handles Unicode, so
// it can process text in all the world’s languages.
func main() {
	// fmt.Println('Hello world')  // single quotes can't be used

	// double Quote Strings - - single line strings
	fmt.Println("Hello go world")
	fmt.Println("Hello\tgo\nworld")
	fmt.Println("Quotes: \"\"")
	fmt.Println("Backslash: \\")

	// fmt.Println("Hello
	// 			 go world")    

	// back-tick strings - Both Single & Multi-line Strings
	fmt.Println(`Hello go world`)
	fmt.Println(`Hello\tgo\nworld`) // Escape sequences were not recognized
	fmt.Println(`Quotes: ""`)
	fmt.Println(`Backslash: \\`)
	fmt.Println(`Hello 
                      world`)
}
