package main

import "fmt"

func main() {
	myStr := "Hello World" // strings are immutable
	fmt.Println("myStr =", myStr)

	// mystr[0] = "F" // cannot assign to mystr[0]

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)

}
