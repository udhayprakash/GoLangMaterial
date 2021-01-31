package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func ReverseUsingLoop(s string) string {
	var revStr string
	for i := len(s) - 1; i >= 0; i-- {
		revStr += string(s[i])
	}
	return revStr
}

func ReverseUsingRange(s string) string {
	var revStr string
	strLen := len(s)
	for i := range s {
		revStr += string(s[strLen-1-i])
	}
	return revStr
}

// Reverse returns a string with the bytes of s in reverse order.
func ReverseByBytes(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

func ReverseUsingRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func ReverseUsingClosure(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

func ReverseUsingRecursion(s string) string {
	if s == "" {
		return ""
	}
	return s[0:0] + ReverseUsingRecursion(s[1:])
}

func ReverseUsingRuneAppend(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}
func main() {
	s := "Hello, world"
	s = "Hello, 世界"
	fmt.Println(s)

	fmt.Println(ReverseUsingLoop(s))
	fmt.Println(ReverseUsingRange(s))
	fmt.Println(ReverseByBytes(s))
	fmt.Println(ReverseUsingRunes(s))
	fmt.Println(Reverse(s))
	fmt.Println(ReverseUsingClosure(s))
	fmt.Println(ReverseUsingRecursion(s))
	fmt.Println(ReverseUsingRuneAppend(s))

}

/*
CASE 1: ASCII chars
	Hello, world
	dlrow ,olleH
	dlrow ,olleH
	dlrow ,olleH
	dlrow ,olleH

CASE 2: utf 8 chars
	Hello, 世界
	界世 ,olleH
	ç¸ä ,olleH
	ç¸ä ,l
	��疸� ,olleH
*/
