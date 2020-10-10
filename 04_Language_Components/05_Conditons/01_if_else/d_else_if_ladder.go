package main

import "fmt"

func main() {
	// A statement can precede conditionals; any variables
	// declared in this statement are available in all
	// branches.
	if num := 21; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else if num > 10 {
		fmt.Println(num, "has multiple digits")
	}
	// else block is optional4
	// parentheses around conditions are not needed, but braces for blocks are needed
}
