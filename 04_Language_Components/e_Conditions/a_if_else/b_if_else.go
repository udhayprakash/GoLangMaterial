package main

import "fmt"

func main() {

	var choice bool // default is false
	if choice == true {
		fmt.Println("Default value for bool is true")
	} else {
		fmt.Println("Default value for bool is false")
	}

	const num1 = 33444
	if num1%2 == 0 {
		fmt.Println("num1 is even number")
	} else {
		fmt.Println("num1 is odd number")
	}
}
