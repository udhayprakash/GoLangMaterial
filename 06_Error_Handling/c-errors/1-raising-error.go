package main

import (
	"errors"
	"fmt"
)

func boom() error {
	return errors.New("OM BHEM BHUSHHH!")
}

func main() {
	boom()

	// err := boom()
	// if err != nil {

	if err := boom(); err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("All's Good")
}
