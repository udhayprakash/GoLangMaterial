package main

import "fmt"

// troubleshoot zeroDivisionError panic

func main() {
	fmt.Println(divide(10, 0))
}

func divide(a, b int) int {
	return a / b
}

// b divide
