package main

import "fmt"

/*
Purpose: Function with two arguments and no return value
*/

// Function Definitions
func personDetails(name string, age int) {
	fmt.Printf("%s's age is %d", name, age)
}

func main() {
	// Function Call
	// personDetails() not enough arguments in call to personDetails
	// personDetails("Ken Thompson") not enough arguments in call to personDetails

	personDetails("Ken Thompson", 62)
	// personDetails("Ken Thompson", "62")
	// cannot use "62" (type untyped string) as type int in argument to personDetails

	// personDetails("Ken Thompson", 61, 123)
	// too many arguments in call to personDetails
}
