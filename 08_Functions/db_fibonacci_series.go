package main

import "fmt"

// Fibonacci series - 0, 1, 1, 2, 3, 5, 8, 13, ....
// position           0  1  2  3  4  5  6  7    ...
func fib(n int) int {
	// fmt.Printf("\tfib(%d)\n", n)
	x, y := 0, 1 // unpacking

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x

}

func main() {
	// Function Call
	fmt.Println("fib(1) = ", fib(1))
	fmt.Println("fib(3) = ", fib(3))
	fmt.Println("fib(6) = ", fib(6))
	fmt.Println()

	// Generating first 25 values
	for i := 0; i < 25; i++ {
		// fmt.Println(i, fib(i))
		fmt.Printf("fib(%2d) = %6d\n", i, fib(i))
	}

}


// Assignment: store the computed values in map, 
// 			and use them for subsequent results
// Memoization design pattern
