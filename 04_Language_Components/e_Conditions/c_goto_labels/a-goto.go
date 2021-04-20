package main

import (
	"fmt"
)

func main() {
	fmt.Println("One")
	goto Prakash
	fmt.Println("Two")
	goto Udhay
	fmt.Println("Last statement")

Prakash:
	fmt.Println("You are in Prakash Block")

Udhay:
	fmt.Println("You Entered Udhay Block")

}

// NOTE: The goto labels need to be defined in order of
// their usage.
