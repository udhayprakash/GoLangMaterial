package main

import "fmt"

func main() {
	// string concatenation
	fmt.Println("Hello" + "world")
	fmt.Println("Hello " + "world")
	fmt.Println("Hello " + `world`)

	//fmt.Println('H' + "H")  // (mismatched types untyped rune and untyped string)

	// string repetition
	//fmt.Println("Hello " * 3)  //(mismatched types string and int)

}
