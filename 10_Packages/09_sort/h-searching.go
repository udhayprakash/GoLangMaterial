package main

import (
	"fmt"
	"sort"
)

func main() {
	// integer slice in unsort order
	intSlice := []int{55, 22, 18, 9, 12, 82, 28, 36, 45, 65}
	x := 18
	pos := sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)
		
	// slice need to be sort in ascending order before to use SearchInts
	sort.Ints(intSlice)	// slice sorted
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)
	
	x = 54
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)
	
	x = 99
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)
	
	x = -5
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)
}