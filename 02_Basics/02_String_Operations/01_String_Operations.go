package main

import "fmt"

func main() {
	// fmt.Println('Hello world')  // single quotes can't be used


	// double Quote Strings - - single line strings
	fmt.Println("Hello go world")
	fmt.Println("Hello\tgo\nworld")
	// fmt.Println("Hello
	// 			 go world")    

	// back-tick strings - Both Single & Multi-line Strings
	fmt.Println(`Hello go world`)
	fmt.Println(`Hello\tgo\nworld`)  // Escape sequeces were nott recognized
	fmt.Println(`Hello 
						world`)
}