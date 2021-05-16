package main

import "fmt"

func main() {
	// Switch cases evaluate cases from top to bottom,
	// stopping when a case succeeds.
	fmt.Println("Enter some number:")
	var value int32
	fmt.Scanf("%d", &value)

	switch value {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
	case 3:
		fmt.Println("value is 3")
	default:
		fmt.Println("default case")
	}

}

// NOTE: default is optional
