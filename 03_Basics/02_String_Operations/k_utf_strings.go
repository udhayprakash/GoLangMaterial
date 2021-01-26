package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"unicode/utf8"
)

// go get golang.org/x/text/cases

func main() {
	s := "Hello, 世界"
	fmt.Println("string length:", len(s))                    // "13"
	fmt.Println("rune count   :", utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i:])
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Println("r=", r, "size=", size)
	}

	c := cases.Fold()
	fmt.Printf("%s %v", c, c.String("grüßen"))
}
