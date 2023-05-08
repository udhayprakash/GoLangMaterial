package main

import "fmt"

// four ways to declare a function with two parameters and one result

func add(x int, y int) int {
	z := x + y
	return z
}

func sub(x, y int) (z int) {
	z = x - y
	return // naked return
}

func first(x int, _ int) int { // _ when not using. ex: future enhancements
	return x
}

func zero(int, int) int {
	return 0
}

func main() {
	// type of a function is sometimes called its signature.
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
	fmt.Println()

	fmt.Println("add(1, 2)   =", add(1, 2))
	fmt.Println("sub(1, 2)   =", sub(1, 2))

	fmt.Println("first(1, 2) =", first(1, 2))
	// fmt.Println("first(1) =", first(1))  // not enough arguments in call to first

	fmt.Println("zero(1, 2)   =", zero(1, 2)) // 0
	fmt.Println("zero(4, 2)   =", zero(4, 2)) // 0
	fmt.Println("zero(4, 6)   =", zero(4, 6)) // 0
	fmt.Println()

	// var num1 int
	var f func(int) int   // zero value of function type is nil
	fmt.Printf("%T\n", f) // "func(int) int"
	// f(3)               // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("f == nil :", f == nil) // true

}
