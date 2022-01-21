package main

import "fmt"


// Switch cases evaluate cases from top to bottom,
// stopping when a case succeeds.

func main() {
	fmt.Print("Enter some number:")

	var value int8
	fmt.Scanf("%d", &value)

	switch value {
	case 1: // same as checking if value == 1
		fmt.Println("value is 1")
	case 2: // value == 2
		fmt.Println("value is 2")
	case 3: // value == 3
		fmt.Println("value is 3")
	default: // similar to else
		fmt.Println("default case")
	}
}

// NOTE: default case is optional
