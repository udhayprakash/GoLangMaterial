package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	x := "my new text is this long"
	y := strings.Repeat("#", utf8.RuneCountInString(x))
	fmt.Println(x) // my new text is this long
	fmt.Println(y) // ########################
}