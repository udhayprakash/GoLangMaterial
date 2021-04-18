package main

import (
	"fmt"
	"strings"
)

func main() {
	var pieces []string

	pieces = strings.Split("", "")
	fmt.Println(pieces, len(pieces))
	fmt.Printf("%v %d\n", pieces, len(pieces))
	fmt.Printf("%q %d\n", pieces, len(pieces))
	fmt.Printf("%#v %d\n", pieces, len(pieces))

	fmt.Println("-----")

	pieces = strings.Split("", ",")
	fmt.Println(pieces, len(pieces))
	fmt.Printf("%v %d\n", pieces, len(pieces))
	fmt.Printf("%q %d\n", pieces, len(pieces))
	fmt.Printf("%#v %d\n", pieces, len(pieces))

	fmt.Println("-----")

	sslice := []string{"", "cb", "a b"}
	fmt.Printf("%%v: %v\n", sslice)
	fmt.Printf("%%+v: %+v\n", sslice)
	fmt.Printf("%%q: %q\n", sslice)
	fmt.Printf("%%#v: %#v\n", sslice)
}

// NOTE: "%#v" print verb is intended to be used for unambiguous printing of a type
