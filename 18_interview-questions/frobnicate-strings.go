package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// frobnicate/tweak a string
func frobnicate(input string) string {
	var result []string

	for _, i := range input {
		switch {
		case unicode.IsSpace(i):
			result = append(result, " ")
			// uncomment to frob the digits
		//case unicode.IsNumber(i):
		//	result = append(result, string(rune(i)))
		case utf8.ValidRune(i):
			result = append(result, string(rune(i)^42))
		}

	}

	return strings.Join(result, "")
}

func main() {
	text := "жѳМѭњЂЯёЧВ 一二三 handle alphabets = ขอพฮศโำฐเฦ abc世界你好123 ペツワケザユプルヂザ"
	fmt.Println(text)
	fmt.Println(frobnicate(text))

	fmt.Println("Invertible test:")
	fmt.Println(frobnicate(frobnicate(text)))
}

// Ref: https://www.techopedia.com/definition/27996/frobnicate
