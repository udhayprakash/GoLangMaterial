package main

import (
	"fmt"
)

//  goto  is used to tranfer control to a labelled statement

func main() {
	fmt.Println("One")
	fmt.Println("Two")

	goto Chennai
	goto Hyderabad

	fmt.Println("Last statement")

Hyderabad:
	fmt.Println("Hyderabad")

Chennai:
	fmt.Println("Chennai")
	// goto Hyderabad--- leads to infinite lop

}

// NOTE: The goto labels need to be defined in order of their usage.
