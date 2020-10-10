package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1 // tuple assignment
	for i := 0; i < n; i++ {
		x, y = y, x+y // tuple assignment
	}
	return x
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Printf("fib(%2d) = %3d\n", i, fib(i))
	}
}
