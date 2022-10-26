package main

import "fmt"

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func main() {
	// G	o	l	a	n	g
	// 0	1	2	3	4	5
	fmt.Println(substr("Golang", 2, 5)) // "lang"
	fmt.Println(substr("Golang", 2, 4)) // "lang"
	fmt.Println(substr("Golang", 2, 3)) // "lan"
	fmt.Println(substr("Golang", 2, 2)) // "la"
	fmt.Println(substr("Golang", 2, 1)) // "l"
	fmt.Println(substr("Golang", 2, 0)) // ""
}
