package main

import (
	"fmt"
)

type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

func someError() error {
	return MyError{"Some error occurred"}
}

func main() {
	if err := someError(); err != nil {
		fmt.Println(err)
	}
}
