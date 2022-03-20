package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func main() {
	_, err := capitalize("")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}
	fmt.Println("Success!")
}
