package main

import "fmt"

func myfunc1(word string) {
	fmt.Printf("In myfunc1: %10T %v\n", word, word)
}

func myfunc2(word interface{}) {
	fmt.Printf("In myfunc2: %10T %v \n", word, word)
}

func main() {
	myfunc1("World")
	//myfunc1(2131)

	myfunc2("World")
	myfunc2(2131)
	myfunc2(123.213)
	myfunc2(12 + 3i)
	myfunc2(true)
	myfunc2([3]int{1, 2, 3})
	myfunc2([]int{1, 2, 3})
}
