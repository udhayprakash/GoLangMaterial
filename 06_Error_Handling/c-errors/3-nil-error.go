package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "Booom !!!"
}

func sayHello(message string) (string, error) {
	if message == "" {
		return "", &MyError{} // error
	}
	return message, nil // simulating no error
}

func main() {
	s, err := sayHello("")
	if err != nil {
		fmt.Println("Unexpected error:", err)
		// return
	}
	
	fmt.Println()

	s, err = sayHello("Hello world")
	if errors.Is(err, nil) {
		fmt.Println("No error occurred")
		fmt.Println("The string:", s)
	}
}
