package main

import "fmt"

func myFunc1(word string) {
	fmt.Printf("In myFunc1: %10T %v\n", word, word)
}


func myFunc2(word interface{}) {
	fmt.Printf("In myFunc2: %10T %v\n", word, word)
}


func main(){
	myFunc1("World")
	//myFunc1(123213) // cannot use 123213 (type untyped int) as type string in argument to myFunc1


	myFunc2("World")
	myFunc2(2131)
	myFunc2(123.213)
	myFunc2(12 + 3i)
	myFunc2(true)
	myFunc2([3]int{1, 2, 3})
	myFunc2([]int{1, 2, 3})

}