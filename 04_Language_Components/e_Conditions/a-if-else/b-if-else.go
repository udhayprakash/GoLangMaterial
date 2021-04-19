package main

import (
	"fmt"
)

func main() {
	var choice bool // default value is false

	if choice == true {
		fmt.Println("Value is true")
	} else {
		fmt.Println("Value is false")
	}

	choice = (true == choice)
	if choice == true {
		fmt.Println("Value2 is true")
	} else {
		fmt.Println("Value2 is false")
	}

	const num1 = 33444
	if num1%2 == 0 {
		fmt.Println("num1 is even number")
	} else {
		fmt.Println("num1 is odd number")
	}

}
