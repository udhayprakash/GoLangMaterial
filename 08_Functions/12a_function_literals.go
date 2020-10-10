package main

import (
	"sort"
	"fmt"
)

func main() {

	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people) // Output: [Bob Dave Alice]


	// It could also have been stored in an intermediate variable:
	people1 := []string{"Alice", "Bob", "Dave"}
	less := func(i, j int) bool {
		return len(people1[i]) < len(people1[j])
	}
	sort.Slice(people1, less)
	fmt.Println(people1) // Output: [Bob Dave Alice]

	// NOTE: less function is a closure: it references the
	//      people variable, which is declared outside the
	//     function.


}
