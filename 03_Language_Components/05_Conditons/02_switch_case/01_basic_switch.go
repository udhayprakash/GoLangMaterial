package main

import (
	"fmt"
)

func main() {
	// Here's a basic `switch`.
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
