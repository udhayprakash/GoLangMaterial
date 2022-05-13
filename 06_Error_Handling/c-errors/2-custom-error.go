package main

import (
	"fmt"
)

type MyError struct {
	message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("ERROR:%v", e.message)
}

func someError() error {
	return MyError{"Some Error Occurred!!!"}
}

func main() {
	someError()

	if err := someError(); err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("All's Good")
}
