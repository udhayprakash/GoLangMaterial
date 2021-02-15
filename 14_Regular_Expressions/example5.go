// Go program to illustrate how to replace
// string with the specified regexp
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Creating and initializing a string
	// Using shorthand declaration
	s := "Geeks-for-Geeks-for-Geeks-for-Geeks-gfg"

	// Replacing all the strings
	// Using ReplaceAllString() method
	m := regexp.MustCompile("^(.*?)Geeks(.*)$")
	Str := "${1}GEEKS$2"
	res := m.ReplaceAllString(s, Str)
	fmt.Println(res)

}
