package main

import "fmt"

func main() {
	if true {
		fmt.Println("It is true")
	}

	if false {
		fmt.Println("It is false")
	}

	if !false {
		fmt.Println("It is !false")
	}

	var result int = 12
	if result == 12 { // relational operation
		fmt.Println("result is 12")
	}

	if result != 12 { // false
		fmt.Println("result is NOT 12")
	}

	if !(result != 12) { // true
		fmt.Println("- !(result != 12)")
	}

	if result > 0 && result <= 12 { // Logical(&&) operation
		fmt.Println("result is between 0 and 12 inclusive")
	}

	if result == -12 || result == 12 { // Logical(||) operation
		fmt.Println("result is either -12 or 12")
	}
}
