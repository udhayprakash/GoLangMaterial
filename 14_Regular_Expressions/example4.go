// Go program to illustrate how to
// replace string with the specified regexp
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Replace string with the specified regexp
	// Using ReplaceAllString() method
	m1 := regexp.MustCompile(`x(p*)y`)

	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "B"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy--", "$1"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "$1P"))
	fmt.Println(m1.ReplaceAllString("xy--xpppyxxppxy-", "${1}Q"))

}
