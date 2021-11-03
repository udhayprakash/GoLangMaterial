package main

import "fmt"

// abcdefghijklmnopqrstuvwxyz
// EX: I am in mozambique now
//     L dp lq prcdpeltxh qrz

func main() {
	fmt.Printf("\t rune       : %c \n", 'a') // a
	fmt.Printf("\t rune       : %q \n", 'a') // 'a'
	fmt.Printf("\t UTF8 code  : %v \n", 'a') // 97
	fmt.Println()

	fmt.Printf("\t rune       : %c \n", 'z')   // z
	fmt.Printf("\t rune       : %q \n", 'z')   // 'z'
	fmt.Printf("\t UTF8 code  : %v \n\n", 'z') // 122
	fmt.Println()

	s := "hello world"
	encryptedString := ""

	for i := 0; i < len(s); i++ {
		// fmt.Println(i, string(s[i]), s[i])

		// fmt.Printf("\t rune       : %c \t", s[i])
		// fmt.Printf("\t UTF8 code  : %v \n", s[i])

		// // Caesar Cipher : character + 3
		// fmt.Printf("\t rune       : %c \t", s[i]+3)
		// fmt.Printf("\t UTF8 code  : %v \n\n", s[i]+3)

		encryptedString += string(s[i] + 3)

	}
	fmt.Println("encryptedString=", encryptedString)

}

// Assignment - caesar cipher  decrypting - 3
 