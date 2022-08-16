// Go program to illustrate how to
// find regexp from the given slice
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Finding regexp from
	// the given slice
	// Using Find() method
	m := regexp.MustCompile(`language`)
	res := m.Find([]byte(`I like Go language this language is "+ 
						"very easy to learn and understand`))

	if res == nil {

		fmt.Printf("Found: %q ", res)
	} else {
		fmt.Printf("Found: %q", res)
	}
}
