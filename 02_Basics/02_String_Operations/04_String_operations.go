package main

import "fmt"

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
