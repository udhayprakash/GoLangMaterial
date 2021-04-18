package main 

import (
	"fmt"
)

func main(){
	// String Concatenation
	fmt.Println("Hello" + "world")
	fmt.Println("Hello" + " " + "world")
	fmt.Println("Hello" + " " + `world`)
	// fmt.Println("Hello" + ' ' + `world`)
	// mismatched types untyped string and untyped rune

	//fmt.Println('H' + "H")  // (mismatched types untyped rune and untyped string)

	// string repetition
	//fmt.Println("Hello " * 3)  //(mismatched types string and int)
	fmt.Println('H' * 3) // 216 
}