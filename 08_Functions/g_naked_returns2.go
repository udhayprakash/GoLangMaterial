package main

import "fmt"

// func add(a, b int) sum int {
func add(a, b int) (sum int) {
	sum = a + b
	return
}

// NOTE: naked return signature must
//  be enclosed in paranthesis.

func main() {
	fmt.Println(add(17, 34))
}
