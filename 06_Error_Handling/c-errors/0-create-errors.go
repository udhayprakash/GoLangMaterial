package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := errors.New("Zero Division Error")
	fmt.Println("Error is :", err)

	timeErr := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", timeErr)
}
