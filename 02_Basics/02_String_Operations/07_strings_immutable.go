package main

import "fmt"

func main() {
	mystr := "Hello World" // strings are immutable

	// mystr[0] = "F" // cannot assign to mystr[0]

	fmt.Println("mystr =", mystr)

	// rune value can be updated
	myrune := 'H'
	fmt.Println("before myrune =", myrune)

	myrune = 'W'
	fmt.Println("after   myrune =", myrune)

}
