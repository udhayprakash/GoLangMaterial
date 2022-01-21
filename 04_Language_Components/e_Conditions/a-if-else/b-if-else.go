package main

import "fmt"

func main() {
	var choice bool // zero value is false

	if choice == true {
		fmt.Println("choice is true")
	} else {
		fmt.Println("choice is false")
	}

	if choice { // indirectly you are checing for choice == true
		fmt.Println("choice is true")
	} else {
		fmt.Println("choice is false")
	}

	choice = (true != choice)
	if choice == true {
		fmt.Println("new choice is true")
	} else {
		fmt.Println("new choice is false")
	}

	const num1 = 33444
	if num1%2 == 0 {
		fmt.Printf("%d is even number", num1)
	} else {
		fmt.Printf("%d is odd number", num1)
	}
}
