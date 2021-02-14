package main

import (
	"fmt"
	"sort"
)

func main() {
	// Method 1
	people := []string{"sudeep", "ram", "udhay", "prakash"}

	fmt.Println("Initially    , people:", people)
	sort.Strings(people)
	fmt.Println("After sorting, people:", people)

	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println("After slicing, people:", people)

	// Method 2
	// It could also have been stored in an intermediate variable:
	people1 := []string{"sudeep", "ram", "udhay", "prakash"}
	less := func(i, j int) bool {
		return len(people1[i]) < len(people1[j])
	}
	sort.Slice(people1, less)
	fmt.Println(people1) // Output: [Bob Dave Alice]

	// NOTE: less function is a closure: it references the
	//      people variable, which is declared outside the
	//     function.

}
