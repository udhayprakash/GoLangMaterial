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

	fmt.Println("â", 'â') // â 226

	fmt.Println(`\t`, "\\t", '\t')
	fmt.Println(`\n`, "\\n", '\n')
	fmt.Println(`\`, "\\", '\\')
	fmt.Println()

	fmt.Println("var/user/somebody")
	fmt.Println("C:\newdir\newfile")
	fmt.Println("C:\\newdir\\newfile")
	fmt.Println(`C:\newdir\newfile`)
	fmt.Println()

	name := "Golang"
	fmt.Println("name = ", name)
	fmt.Printf("name = %s\n", name)

	// Runes are printed with %c, or with %q if quoting is desired:
	ascii := 'a'
	fmt.Printf("%d %c %q\n", ascii, ascii, ascii) // 97 a 'a'
	fmt.Printf("%d %c %q\n", ascii)               // 97 %!c(MISSING) %!q(MISSING)
	fmt.Printf("%d %[1]c %[1]q\n", ascii)         // 97 a 'a'

	unicode := 'D'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 68 D 'D'

	unicode = 'न'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 2344 न 'न'

	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", newline) // 10  	'\n'
	fmt.Printf(`%d %[1]c %[1]q\n`, newline) // 10  	'\n'
}
