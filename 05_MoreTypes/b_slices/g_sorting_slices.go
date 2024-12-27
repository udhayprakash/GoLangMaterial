package main

import (
	"fmt"
	"sort"
)

// Arrays cant be sorted; only the slices can be sorted
func main() {
	// === sorting strings
	strs := []string{"c", "a", "b"}
	fmt.Printf("\nstrs = %T", strs)                  // []string
	fmt.Printf("\nBefore sorting - strs = %v", strs) // [c a b]

	// sorting the slices
	sort.Strings(strs)
	fmt.Printf("\nAfter sorting -  strs = %v", strs)
	fmt.Println(sort.StringsAreSorted(strs)) // true

	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Printf("\nAfter rev sorting-strs = %v\n", strs) // [c b a]

	strs = []string{"c", "a", "b"}
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Printf("\nAfter rev sorting-strs = %v\n", strs) // [c b a]
	fmt.Println(sort.StringsAreSorted(strs))            // false
	// By default, in ascending order, it will check
	// Assignment: how to check sorted in descending order

	fmt.Println()

	// === sorting integers
	ints := []int{213, 234, 213, 43, 2, 4}
	fmt.Println("\nBefore sorting - ints = ", ints)
	sort.Ints(ints)
	fmt.Println("After  sorting - ints = ", ints) // [2 4 43 213 213 234]

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s) // true

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("After rev sorting - ints = ", ints) //  [234 213 213 43 4 2]

	// === sorting floats
	fltSlice := []float64{18787677.878716, 565435.321, 7888.545, 8787677.8716, 987654.252} // unsorted
	fmt.Println("Slice BEFORE sort: ", fltSlice)

	sort.Float64s(fltSlice)
	fmt.Println("Slice AFTER sort: ", fltSlice)
	fmt.Println(sort.Float64sAreSorted(fltSlice)) // true

}

// Assignment: Explore the sort module,
// and try to sorting in descending order for both strings and ints
