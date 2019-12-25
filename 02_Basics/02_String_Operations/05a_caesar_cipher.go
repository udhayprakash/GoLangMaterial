package main

import "fmt"

func main() {
	fmt.Printf("\t rune       : %c \n", 'a') // a
	fmt.Printf("\t UTF8 code  : %v \n", 'a') // 97

	fmt.Printf("\t rune       : %c \n", 'z') // z
	fmt.Printf("\t UTF8 code  : %v \n\n", 'z') // 122


	s := "hello worldz"
	for i := 0; i < len(s); i++ {
		fmt.Printf("\t rune       : %c \n", s[i])
		fmt.Printf("\t UTF8 code  : %v \n", s[i])

		// Caesar Cipher : character + 3
		fmt.Printf("\t rune       : %c \n", s[i]+3)
		fmt.Printf("\t UTF8 code  : %v \n\n", s[i]+3)

	}
}
