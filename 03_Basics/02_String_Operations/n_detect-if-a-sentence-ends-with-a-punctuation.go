package main

import (
	"log"
	"unicode"
)

func main() {
	//str := "this is a line with fullstop."
	str := "this is a line with without fullstop"

	lastChar := str[len(str)-1:]
	lastCharRune := []rune(lastChar)
	if unicode.IsPunct(lastCharRune[0]) { // this is how to convert string to rune ;-)
		log.Println("lastChar is a punctuation :", lastChar)
	} else {
		log.Println("lastChar is not a punctuation :", lastChar)
	}
}
