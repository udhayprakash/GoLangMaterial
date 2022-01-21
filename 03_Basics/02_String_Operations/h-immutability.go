package main

import "fmt"

/*
mutable object -- one which supports chnges in values, in place
immutable object - one which doesnt support inplace changes
	ex: strings

*/

func main() {
	myStr := "Hello world"
	fmt.Println("myStr = ", myStr, &myStr)
	fmt.Printf("myStr = %#v\n", myStr)

	// Indexing
	fmt.Println("myStr[0]= ", myStr[0], string(myStr[0])) //  72  H

	// myStr[0] = "F"
	// cannot assign to myStr[0] (strings are immutable)

	// myStr[0] = 'F'
	// cannot assign to myStr[0] (strings are immutable)

	// Slicing
	fmt.Println("myStr[6:11] = ", myStr[6:11]) //  world

	// myStr[6:11] = "Bould"
	// annot assign to myStr[6:11] (strings are immutable)

	// Overwriting
	myStr = "Fello world"
	fmt.Println("After OVERWRITE: myStr=", myStr, &myStr)

	myStr = myStr[:6] + "Bould"
	fmt.Println("After OVERWRITE: myStr=", myStr, &myStr)
	// NOTE: strings are immutable; runes are mutable
	fmt.Println()

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)
}
