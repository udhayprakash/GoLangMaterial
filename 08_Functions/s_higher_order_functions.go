package main

import (
	"fmt"
	"strings"
)

/*
Higher-Order Functions
	- Functions which were designed to work on top of other functions
	- strings.Map
		- applies a function to each character of a string,
		- joining the results to make another string
*/

func add2(character rune) rune {
	return character + 1
}

func caesarCipher(character rune) rune {
	if character == ' ' {
		return character
	}
	return character + 3

}

func main() {
	fmt.Println(strings.Map(add2, "abcdefgh")) // "bcdefghi"
	fmt.Println(strings.Map(add2, "bcdefghi")) // "cdefghij"
	fmt.Println(strings.Map(add2, "Golang"))   // "Hpmboh"
	fmt.Println(strings.Map(add2, "Udhay123")) // "Veibz234"

	fmt.Println(strings.Map(caesarCipher, "Let's attack at the river edge to night"))
	// Ohw*v dwwdfn dw wkh ulyhu hgjh wr qljkw

}
