package main

import "fmt"

func main() {
	// var x2 interface{} = true
	// fmt.Println(x2.(string)) // interface {} is bool, not string

	var x interface{} = "foo"
	fmt.Println(x.(string)) // foo

	var s string = x.(string)
	fmt.Println(s) // "foo"

	s, ok := x.(string)
	fmt.Println(s, ok) // "foo true"

	n, ok := x.(int)
	fmt.Println(n, ok) // "0 false"

	// n = x.(int)
	// panic: interface conversion: interface {} is string, not int

	var y interface{} = "123"
	fmt.Println("y.(string) =", y.(string))

	// fmt.Println("y.(int) =", y.(int))
	// panic: interface conversion: interface {} is string, not int

}
