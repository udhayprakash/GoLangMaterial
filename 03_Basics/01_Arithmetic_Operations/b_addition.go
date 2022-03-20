package main

import "fmt"

func main() {
	fmt.Println("1 + 1 		=", 1+1)        // int + int
	fmt.Println("1.0 + 1.0 	=", 1.0+1.0) // float + float
	fmt.Println("1.0 + 1.1 	=", 1.0+1.1) // float + float

	fmt.Println("1.3 + 222  =", 1.3+222) // float + int

	fmt.Println("2i + 222  =", 2i+222) // complex + int

	// 12 + true
	// (mismatched types untyped int and untyped bool)

	fmt.Println("'a' + 10   =", 'a'+10) // 107

	// "golang" + 'g'
	// (mismatched types untyped string and untyped rune)
}
