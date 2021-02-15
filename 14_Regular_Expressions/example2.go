// Go program to illustrate how to find
// all the regexp from the given slice
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Finding all regexp from
	// the given string
	// Using FindAllString () method
	m := regexp.MustCompile(`geeks.`)

	fmt.Println(m.FindAllString("GeeksgeeksGeeks, geeks", -1))
	fmt.Println(m.FindAllString("Hello! geeksForGEEKsgeeks-geeks", 2))
	fmt.Println(m.FindAllString("I like Go language", 0))
	fmt.Println(m.FindAllString("Hello, Welcome", 1))

}
