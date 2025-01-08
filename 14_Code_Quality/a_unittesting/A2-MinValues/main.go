package main

import (
	"errors"
	"fmt"
)

// Div performs integer division and returns an error if the denominator is zero.
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := Div(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // Output: Result: 5
	}

	result, err = Div(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero
	} else {
		fmt.Println("Result:", result)
	}
}
