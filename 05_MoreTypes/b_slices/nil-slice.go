package main

import "fmt"

func main() {
	var emptySlice []int
	fmt.Println("emptySlice       :", emptySlice)
	fmt.Println("emptySlice == nil:", emptySlice == nil)
	fmt.Println("len(emptySlice)  :", len(emptySlice))
	if emptySlice == nil {
		fmt.Println("Empty slice")
	}
	fmt.Println()

	// Remove all elements in a slice
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a)) // [] 0 0

	// Keep allocated memory
	a = []string{"A", "B", "C", "D", "E"}
	a = a[:0]
	fmt.Println(a, len(a), cap(a)) // [] 0 5
	// If the slice is extended again, the original data reappears.
	fmt.Println(a[:2]) // [A B]
}
