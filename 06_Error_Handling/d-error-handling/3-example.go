package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Denominator should not be zero")
	}
	return x / y, nil

}

func main() {
	// fmt.Println("div(30, 2)=", div(30, 2))
	// multiple-value div(30, 2) (value of type (int, error)) in single-value context

	divResult, err := div(30, 2)
	if err != nil {
		fmt.Printf("Error is %v\n", err)
	} else {
		fmt.Println("divResult = ", divResult)
	}

	divResult, err = div(30, 0)
	if err != nil {
		fmt.Printf("Error is %v\n", err)
	} else {
		fmt.Println("divResult = ", divResult)
	}

}
