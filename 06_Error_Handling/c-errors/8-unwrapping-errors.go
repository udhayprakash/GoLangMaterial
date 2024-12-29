package main

import (
	"errors"
	"fmt"
)

var rootError = errors.New("root error")

func CauseAnError() error {
	return fmt.Errorf("wrapping an error: %w", rootError)
}

func main() {
	err := CauseAnError()
	if errors.Unwrap(err) == rootError {
		fmt.Println("Unwrapped to the root error:", rootError)
	}
}
