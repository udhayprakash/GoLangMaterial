package main

import (
	"errors"
	"fmt"
)

func readFile() error {
	return errors.New("file not found")
}

func processFile() error {
	err := readFile()
	if err != nil {
		// Wrapping error with context
		return fmt.Errorf("processing failed: %w", err)
	}
	return nil
}

func main() {
	err := processFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
