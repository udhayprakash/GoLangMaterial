package main

import (
	"fmt"
)

type MyError struct {
	message string
	code    int
}

func (e MyError) Error() string {
	return fmt.Sprintf("ERROR:%v (Code: %d)", e.message, e.code)
}

func someError() error {
	return MyError{"Some Error Occurred!!!", 1001}
}

func main() {
	if err := someError(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
