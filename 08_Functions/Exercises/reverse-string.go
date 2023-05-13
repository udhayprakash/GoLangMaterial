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

func ReverseUsingUTF8Decode(s string) (rs string) {
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		w = width
		rs = string(r) + rs
	}
	return
}

// function to reverse given arr without generating new one
// with the help of left and right pointers
func reverse(start, end int, arr []int) {
	for start < end {

		// swapping both indexes
		// using left and right pointers
		arr[start], arr[end] = arr[end], arr[start]

		// moving forward
		start++
		// moving backwards
		end--
	}
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
	fmt.Println(ReverseUsingUTF8Decode(s))

	// non-reversed array
	arr := []int{1, 2, 3, 4, 5, 6}

	// left pointer
	start := 0

	// right pointer
	end := len(arr) - 1

	fmt.Printf("non reversed array: %d\n", arr)

	// calling reverse function
	reverse(start, end, arr)

	fmt.Printf("reversed array: %d\n", arr)
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
