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
		fmt.Println("------------ONE---------------")
	case 2:
		fmt.Println("------------TWO---------------")
	case 3:
		fmt.Println("------------THREE-------------")
	default:
		fmt.Println("__INVALID NUMBER__")
	}
	// NOTE: default is optional
}
