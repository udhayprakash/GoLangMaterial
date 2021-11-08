package main

import "fmt"

// four ways to declare a function with
//        two parameters and one result
func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
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

	// zero value of function type is nil
	var f func(int) int
	// f(3) // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("f == nil :", f == nil)

}
 