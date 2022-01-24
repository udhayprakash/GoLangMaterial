package main

import (
	"fmt"
	"sort"
)

// Arrays cant be sorted; only the slices can be sorted
func main() {

	strs := []string{"c", "a", "b"}
	fmt.Printf("\nstrs = %T", strs)                  // []string
	fmt.Printf("\nBefore sorting - strs = %v", strs) // [c a b]

	// sorting the slices
	sort.Strings(strs)
	fmt.Printf("\nAfter sorting -  strs = %v", strs)

	ints := []int{213, 234, 213, 43, 2, 4}
	fmt.Println("\nBefore sorting - ints = ", ints)
	sort.Ints(ints)
	fmt.Println("After  sorting - ints = ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s) // true

}

// Assignment: Explore the sort module,
// and try to sorting in descending order for both strings and ints
