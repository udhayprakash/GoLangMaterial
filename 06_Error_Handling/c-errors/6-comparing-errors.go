package main

import (
	"errors"
	"fmt"
)

var errOne = errors.New("first error")
var errTwo = errors.New("second error")

func DoSomething() error {
	return errors.Join(errOne, errTwo)
}

func main() {
	err := DoSomething()
	if errors.Is(err, errOne) {
		fmt.Println("The error includes the first error.")
	}
	if errors.Is(err, errTwo) {
		fmt.Println("The error includes the second error.")
	}
	fmt.Println("Joined Error:", err)
}
