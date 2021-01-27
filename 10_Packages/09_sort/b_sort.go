package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1, 6, 7}
	fmt.Println("Initially, values         =", values)
	fmt.Println("sort.IntsAreSorted(values)=", sort.IntsAreSorted(values))

	sort.Ints(values)
	fmt.Println("\nAfter sort, values        =", values)
	fmt.Println("sort.IntsAreSorted(values)=", sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"

}
