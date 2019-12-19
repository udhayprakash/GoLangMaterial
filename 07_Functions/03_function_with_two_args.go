package main

import "fmt"

/*
Purpose: Function with two arguments and no return value

*/

// Function Definition
func personDetails(name string, age int) {
	fmt.Printf("%s's age is %d", name, age)
}

func main() {

	// Function call
	// personDetails()  // not enough arguments in call to personDetails
	// personDetails("Ken Thompson") // not enough arguments in call to personDetails

	personDetails("Ken Thompson", 61)
	// personDetails("Ken Thompson", 61, 123) // too many arguments in call to personDetails

}
