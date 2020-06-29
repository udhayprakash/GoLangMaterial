package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Printf("len(s) =%d\n\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("\t data type  : %T \n", s[i])

		fmt.Printf("\t rune       : %c \n", s[i])
		fmt.Printf("\t string     : %s \n", s[i])
		fmt.Printf("\t quoted     : %q \n", s[i])
		fmt.Printf("\t UTF8 code  : %v \n", s[i])

		fmt.Printf("\t Hexadecimal: %x \n", s[i])
		fmt.Printf("\t Octal      : %o \n", s[i])
		fmt.Printf("\t Binary     : %b \n\n", s[i])
	}

}
