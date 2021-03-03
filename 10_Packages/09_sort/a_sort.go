package main

import (
	"fmt"
	"sort"
)

var strSlice sort.StringSlice = []string{"apple", "durian", "kiwi", "banana"}

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	fmt.Println("Original : ", strSlice[:])

	strSlice.Sort()

	fmt.Println("Sort : ", strSlice[:])

	sort.Sort(sort.Reverse(strSlice[:]))

	fmt.Println("Reverse : ", strSlice[:])

}

// Note: sorting is in-place, so it changes the given slice and
// doesn't return a new one
