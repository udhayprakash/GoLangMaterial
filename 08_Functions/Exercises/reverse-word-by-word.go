package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
NOTE:
	strings.Split() splits based on given substring
	strings.Fields() splits based on tab, newline, formfeed,
            vertical tab, carriage return, space, and
            the Unicode characters for next line and space
*/

// reverseWords Splits a string on whitespace and reverses
// each of the individual words
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// reverseWordsFunc Reverses a string word by word,
// based on a custom function
func reverseWordsFunc(s string) string {
	words := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("one two, three!"))
	fmt.Println(reverseWordsFunc("one two, three!"))
}
