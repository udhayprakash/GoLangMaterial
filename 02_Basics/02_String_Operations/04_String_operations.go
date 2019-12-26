package main

import "fmt"

// UTF-8 is the encoding scheme in Go
/*
Escape Sequences:
	\a "alert" or bell
	\b backspace
	\f form feed
	\n newline
	\r carriage return
	\t tab
	\v vertical tab
	\' single quote (only in the rune literal '\'')
	\" double quote (only within "..." literals)
	\\ backslash

*/
func main() {
	// Strings are represented with double quotes or back-tick
	// runes are represented with single quotes
	fmt.Println("a") // output: a
	fmt.Println('a') // output: 97

	fmt.Println("Җ", 'Җ')
	fmt.Println("\t", '\t')
	fmt.Println("\n", '\n')
	fmt.Println("\\", '\\')

	// Runes are printed with %c, or with %q if quoting is desired:
	ascii := 'a'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"

}
