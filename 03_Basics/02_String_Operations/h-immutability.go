package main

import "fmt"

func main() {
	myStr := "Hello world"
	fmt.Println("myStr= ", myStr, &myStr)

	// myStr[0] = "F"
	// cannot assign to myStr[0] (strings are immutable)

	fmt.Println("myStr[6:11] = ", myStr[6:11])

	// myStr[6:11] = "Bold"
	// cannot assign to myStr[6:11] (strings are immutable)

	myStr = myStr[:6] + "Bold"
	fmt.Println("After OVERWRITE: myStr=", myStr, &myStr)

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)

}

/*
NOTE: strings are immutable; runes are mutable
*/
