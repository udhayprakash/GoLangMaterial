// Go program to illustrate how to find
// all the regexp from the given slice
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Finding the number from
	// the given string
	// Using FindAllStrings() method
	s := "I45, like345, Go-234 langu34age"

	m := regexp.MustCompile(`[-]?\d[\d]*[\]?[\d{2}]*`)
	res := m.FindAllString(s, 2)
	for _, ele := range res {
		fmt.Println("Number:", ele)
	}
}
