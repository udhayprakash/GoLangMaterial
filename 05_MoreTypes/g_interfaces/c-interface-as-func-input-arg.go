package main

import "fmt"

func myFunc1(word string) {
	fmt.Printf("In myFunc1: %10T %v\n", word, word)
}

func myFunc2(word interface{}) {
	fmt.Printf("In myFunc2: %10T %v\n", word, word)
}

func main() {
	myFunc1("World")
	myFunc1("123112")
	// myFunc1(123112)
	// cannot use 123112 (type untyped int) as type string in argument to myFunc1
	fmt.Println()

	myFunc2("World")         // string
	myFunc2("2131")          // string
	myFunc2(2131)            // int
	myFunc2('2')             // rune
	myFunc2(123.213)         // float
	myFunc2(12 + 3i)         // complex
	myFunc2(true)            // bool
	myFunc2([3]int{1, 2, 3}) // array
	myFunc2([]int{1, 2, 3})  // slice

}
