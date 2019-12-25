package main

import "fmt"

// UTF-8 is the encoding scheme in Go
func main() {
	// Strings are represented with double quotes or back-tick
	// runes are represented with single quotes
	fmt.Println("a") // output: a
	fmt.Println('a') // output: 97

	fmt.Println("Җ", 'Җ')
	fmt.Println("\t", '\t')
	fmt.Println("\n", '\n')
	fmt.Println("\\", '\\')
}
