package main

import (
	"fmt"
	"time"
)

var language string

func main() {
	go func() {
		language = "Golang"
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("language:", language)
}

// Goroutines are not guaranteed to happen
// before any event in the program.
