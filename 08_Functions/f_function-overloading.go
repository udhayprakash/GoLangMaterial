package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// When two or more consecutive named function parameters
// share a type, you can omit the type from all but the last.

func add(x, y int) int {
	return x + y
}

// func add(x, y, z int) int {
// 	return x + y + z
// }

func addThree(p, q int, r float32) int {
	return p + q + int(r)
}

// NOTE: Golang supports neither function-overloading, nor function-overwriting

func main() {
	fmt.Println("add(42, 13)=", add(42, 13))

	fmt.Println("addThree(23, 45, 5.6):", addThree(23, 45, 5.6))
}
