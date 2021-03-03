package main

import "fmt"

func main() {
	fmt.Printf("add: %d\n", add(2, 6))
	fmt.Printf("sub: %d\n", sub(2, 6))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}
