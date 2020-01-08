package main

import (
	"errors"
	"fmt"
)

func boom() error {
	return errors.New("barnacles")
}

func main() {
	if err := boom(); err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Anchors away!")
}