package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("*b=", *b)

	fmt.Println("Before ptr change - a=", a, " b=", b, " *b=", *b)
	*b = 300
	fmt.Println("After ptr change  - a=", a, " b=", b, " *b=", *b)

	// NOTE - pointer arithmetics is not supported in go; But can be done using "unsafe" package
}
