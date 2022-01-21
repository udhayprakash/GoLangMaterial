package main

import (
	"fmt"
	"strings"
)

func main() {

	var pieces []string // slice data type

	pieces = strings.Split("", "")
	fmt.Println(pieces, len(pieces))            // [] 0
	fmt.Printf("%v %d\n", pieces, len(pieces))  // [] 0
	fmt.Printf("%q %d\n", pieces, len(pieces))  // [] 0
	fmt.Printf("%#v %d\n", pieces, len(pieces)) // []string{} 0

	fmt.Println("-----")

	pieces = strings.Split("", ",")
	fmt.Println(pieces, len(pieces))            // [] 1
	fmt.Printf("%v %d\n", pieces, len(pieces))  // [] 1
	fmt.Printf("%q %d\n", pieces, len(pieces))  // [""] 1
	fmt.Printf("%#v %d\n", pieces, len(pieces)) // ]string{""} 1

	fmt.Println("-----")

	sslice := []string{"", "cb", "a b"}
	fmt.Printf("%%v: %v\n", sslice)   // %v: [ cb a b]
	fmt.Printf("%%+v: %+v\n", sslice) // %+v: [ cb a b]
	fmt.Printf("%%q: %q\n", sslice)   // %q: ["" "cb" "a b"]
	fmt.Printf("%%#v: %#v\n", sslice) // %#v: []string{"", "cb", "a b"}
}

// NOTE: "%#v" print verb is intended to be used for unambiguous printing of a type
