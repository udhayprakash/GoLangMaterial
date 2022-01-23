package main

import (
	"fmt"
	"sort"
)

// Arrays cant be sorted; only the slices can be sorted
func main() {

	strs := []string{"c", "a", "b"}
	fmt.Printf("\nstrs = %T", strs)
	fmt.Printf("\nBefore sorting - strs = %v", strs)

	// sorting the slices
	sort.Strings(strs)
	fmt.Printf("\nAfter sorting -  strs = %v", strs)

	ints := []int{213, 234, 213, 43, 2, 4}
	fmt.Println("\nBefore sorting - ints = ", ints)
	sort.Ints(ints)
	fmt.Println("After  sorting - ints = ", ints)
	fmt.Println()

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

}

// Assignment: Explore the sort module
