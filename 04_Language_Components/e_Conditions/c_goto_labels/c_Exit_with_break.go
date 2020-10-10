package main

import (
	"fmt"
)
// break statement terminates execution of the
// innermost for, switch, or select statement.

// If we need to break out of a surrounding loop,
// not the switch, you can put a label on the loop
// and break to that label.
func main() {
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}
