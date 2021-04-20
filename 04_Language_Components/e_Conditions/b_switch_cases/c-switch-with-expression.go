package main

import "fmt"

// `switch with expression`.
func main() {
	fmt.Println("Enter some number:")
	var value int
	fmt.Scanf("%d", &value)

	fmt.Println("value=", value)

	switch value % 2 {
	case 0:
		fmt.Println("It is Even number")
	case 1:
		fmt.Println("It is Odd number")
	}

	// Default is optional

}
