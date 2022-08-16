// Go program to illustrate how to
// find regexp from the given slice
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Finding regexp from the
	// given slice of bytes
	// Using Find() method
	m := regexp.MustCompile(`geeks?`)

	fmt.Printf("%q\n", m.Find([]byte(`GeeksgeeksGeeks, geeks`)))
	fmt.Printf("%q\n", m.Find([]byte(`Hello! geeksForGEEKs`)))
	fmt.Printf("%q\n", m.Find([]byte(`I like Go language`)))
	fmt.Printf("%q\n", m.Find([]byte(`Hello, Welcome`)))

}
