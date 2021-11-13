package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints   :", ints)

	// To check the sorting order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted :", s)

	var strSlice sort.StringSlice = []string{"apple", "durian", "kiwi", "banana"}
	fmt.Println("Original :", strSlice)
	fmt.Println("Original :", strSlice[:])

	strSlice.Sort()
	fmt.Println("Sort     :", strSlice[:])

	sort.Sort(sort.Reverse(strSlice[:]))
	fmt.Println("Reverse  :", strSlice[:])
}

// Note: sorting is in-place, so it changes the given slice and
// doesn't return a new one
