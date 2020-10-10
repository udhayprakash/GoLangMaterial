package main

import "fmt"

// array's length is part of its type, so arrays cannot be resized.
func main() {
	var a [2]string
	fmt.Println("before inserting values: a = ", a)
	// Inserting Values
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])  // Array Indexing
	fmt.Println("after inserting values: a = ", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes  =", primes)  // [2 3 5 7 11 13]

	primes1 := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes1 =", primes1) // [2 3 5 7 11 13 0 0 0 0]
	// Observe zero-values

	// length of sequence - It is giving initializer array size
	fmt.Println("len(primes)  = ", len(primes))
	fmt.Println("len(primes1) = ", len(primes1))

	// 2d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
