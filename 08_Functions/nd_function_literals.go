package main

import (
	"fmt"
	"sort"
)

func main() {
	// Method 1
	people := []string{"john", "shyam", "ashok", "prakash"}

	fmt.Println("Initially    , people:", people)
	sort.Strings(people) // alphanumeric sorting
	fmt.Println("After sorting, people:", people)

	// Custom sorting := sorting by length
	sort.Slice(people, func(p1Pos, p2Pos int) bool {
		return len(people[p1Pos]) < len(people[p2Pos])
	})
	fmt.Println("After slicing, people:", people)

	// Method 2
	// It could also have been stored in an intermediate variable:
	people1 := []string{"sudeep", "ram", "udhay", "prakash"}
	lesserLength := func(i, j int) bool {
		return len(people1[i]) < len(people1[j])
	}
	sort.Slice(people1, lesserLength)
	fmt.Println(people1) // Output: [Bob Dave Alice]

	// NOTE: lesserLength function is a closure: it references the
	//      people variable, which is declared outside the
	//     function.

}
