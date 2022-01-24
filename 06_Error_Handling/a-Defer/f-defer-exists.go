package main

import (
	"fmt"
	"os"
)

func main() {
	statusCode := 0

	defer func() {
		fmt.Println("Defer never reached")
	}()

	// os.Exit(statusCode)  // on occurrance, even registered defers wont execute

	defer func() {
		os.Exit(statusCode)
	}()

	defer func() {
		fmt.Println("LIFO - this defer executes first")
	}()

	statusCode = 1
}

// OUTPUT:
// exit status 1

// Exits with a given status:
//    Status code 0 means success
//    Non-zero status code means an error
