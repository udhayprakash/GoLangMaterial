package main

import "fmt"

func main() {
	myStr := "Hello world"
	fmt.Println("myStr= ", myStr, &myStr)

	// Indexing
	fmt.Println("myStr[0]= ", myStr[0], string(myStr[0])) //  72  H

	// myStr[0] = "F"
	// error: invalid left hand side of assignment
	// 	// cannot assign to myStr[0] (strings are immutable)

	// Slicing
	fmt.Println("myStr[6:11] = ", myStr[6:11])

	// myStr[6:11] = "Bould"
	// error: invalid left hand side of assignment
	// 	// cannot assign to myStr[6:11] (strings are immutable)

	// Overwriting
	myStr = myStr[:6] + "Bould"
	fmt.Println("After OVERWRITE: myStr=", myStr, &myStr)

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)
}
 