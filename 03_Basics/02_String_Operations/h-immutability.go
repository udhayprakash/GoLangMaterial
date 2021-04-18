package main

import "fmt"

func main() {
	myStr := "Hello world"
	fmt.Println("myStr= ", myStr)

	// myStr[0] = "F" // cannot assign to myStr[0]
	fmt.Println("myStr[6:11] = ", myStr[6:11])
	// myStr[6:11] = "Bold" // cannot assign to myStr[6:11]

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)

}

/*
NOTE:
	1. strings are immutable
*/
