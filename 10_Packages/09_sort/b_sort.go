package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1, 6, 7, 8, 2}
	fmt.Println("Initially, values         =", values)
	fmt.Println("sort.IntsAreSorted(values)=", sort.IntsAreSorted(values))

	sort.Ints(values)
	fmt.Println("\nAfter sort, values        =", values)
	fmt.Println("sort.IntsAreSorted(values)=", sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"


	// indexing a value in slice
	index := sort.SearchInts(values, 7)
  	fmt.Println("7 is at index:", index)

	index = sort.Search(len(values), func(i int) bool {
	return values[i] >= 7
	})
	fmt.Println("The first number >= 7 is at index:", index)
	// fmt.Println("The first number >= 7 is:", values[index])
}
