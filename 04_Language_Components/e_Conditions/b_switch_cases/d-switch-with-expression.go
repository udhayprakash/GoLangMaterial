package main

import "fmt"

// `switch with expression`.
func main() {
	fmt.Print("Enter some number:")

	var value int
	fmt.Scanf("%d", &value)
	fmt.Println("value=", value, "value % 2=", value%2)

	if value%2 == 0 {
		fmt.Println("It is EVEN number")
	} else if value%2 == 1 {
		fmt.Println("It is ODD number")
	}

	switch value % 2 {
	case 0:
		fmt.Println("It is EVEN number")
	case 1:
		fmt.Println("It is ODD number")
	}

}
