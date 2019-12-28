package main

import (
	"fmt"
	"net/url"
)

func main() {
	// = true
	fmt.Println(isValidUrl("http://www.golangcode.com"))

	// = false
	fmt.Println(isValidUrl("golangcode.com"))

	// = false
	fmt.Println(isValidUrl(""))
}

// isValidUrl tests a string to determine if it is a well-structured url or not.
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}
