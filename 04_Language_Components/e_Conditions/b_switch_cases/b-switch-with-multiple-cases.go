package main

import "fmt"

// `switch with Multiple cases`.
func main() {
	fmt.Print("Enter some number (0-9):")

	var value int
	fmt.Scanf("%d", &value)

	switch value {
	case 0:
		fmt.Println("It is zero")
	case 1, 3, 5, 7, 9: //  value == 1 || value == 3 || value == 5 || ...
		fmt.Println("It is Odd number")
	case 2, 4, 6, 8:
		fmt.Println("It is Even number")
	default:
		fmt.Println("Out of range")
	}
}
