package main

import (
	"fmt"
)

func main() {
	var x interface{} = "foo"
	var s string = x.(string)
	fmt.Println(s) // "foo"

	s, ok := x.(string)
	fmt.Println(s, ok) // "foo true"

	n, ok := x.(int)
	fmt.Println(n, ok) // "0 false"

	n = x.(int) // ILLEGAL
}
