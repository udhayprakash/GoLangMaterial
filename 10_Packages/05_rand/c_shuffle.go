package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Package rand implements a cryptographically secure
// random number generator.

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Initially     : a=", a) // [1 2 3 4 5 6 7 8]

	rand.Seed(111)
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println("After Shuffle : a=", a) // [2 4 8 5 1 6 3 7]

	rand.Seed(222)
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println("After Shuffle : a=", a) // [7 5 6 4 8 1 2 3]

	// Unpredictable
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println("\nAfter Shuffle : a=", a) // unpredictable

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println("After Shuffle : a=", a) // unpredictable

}
