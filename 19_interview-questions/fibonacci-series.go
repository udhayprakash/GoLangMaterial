package main

import "fmt"

// recursive
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// Iterative
func fibIterative() func() int {
	n1, n2 := 0, 1
	return func() int {
		r := n1
		n1, n2 = n2, n1+n2
		return r
	}
}

func main() {
	n := 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibRecursive(i))
	}

	fmt.Println()
	fibIter := fibIterative()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibIter())
	}
	fmt.Println()
}
