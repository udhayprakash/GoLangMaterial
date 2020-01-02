package main

import (
	"fmt"
)

var language string

func main() {
	go func() {
		language = "Golang"
	}()

	fmt.Println("language:", language)
}

// Goroutines are not guaranteed to happen
// before any event in the program.
