package main

import (
	"fmt"
)

func main() {
	fmt.Println("One")
	fmt.Println("Two")
	goto India

	fmt.Println("Last statement")
	goto Mozambique

Mozambique:
	fmt.Println("I am in Mozambique block ")

India:
	fmt.Println("I am in India block ")
	// goto Mozambique --- leads to infinite lop

}

// NOTE: The goto labels need to be defined in order of
// their usage. 