package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main(){
	x := "my new text is this long"
	fmt.Println("x=", x)

	y := strings.Repeat("#", 10)
	fmt.Println("y=", y)

	z := strings.Repeat("#", len(x))
	fmt.Println("z=", z)

	p := strings.Repeat("#", utf8.RuneCountInString(x))
	fmt.Println("p=", p)

}
