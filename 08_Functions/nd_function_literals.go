package main

import (
	"fmt"
	"sort"
)

func main() {
	// Method 1
	people := []string{"john", "shyam", "ashok", "prakash"}

	fmt.Println("Initially    , people:", people) // [john shyam ashok prakash]
	sort.Strings(people)                          // alphanumeric sorting
	fmt.Println("After sorting, people:", people) // [ashok john prakash shyam]

	// Custom sorting := sorting by length -- ascending order
	sort.Slice(people, func(pos1, pos2 int) bool {
		return len(people[pos1]) < len(people[pos2])
	})
	fmt.Println("After slicing, people:", people) // [john ashok shyam prakash]

	// Custom sorting := sorting by length  -- descending order
	sort.Slice(people, func(pos1, pos2 int) bool {
		return len(people[pos1]) > len(people[pos2])
	})
	fmt.Println("After slicing, people:", people) // [prakash ashok shyam john]

	// Method 2
	// It could also have been stored in an intermediate variable:
	people1 := []string{"sudeep", "ram", "udhay", "prakash"}

	lesserLength := func(i, j int) bool {
		return len(people1[i]) < len(people1[j])
	}
	sort.Slice(people1, lesserLength)
	fmt.Println(people1) // [ram udhay sudeep prakash]

	// NOTE: lesserLength function is a closure: it references the
	//      people variable, which is declared outside the function.

	greaterLength := func(i, j int) bool {
		return len(people1[i]) > len(people1[j])
	}
	sort.Slice(people1, greaterLength)
	fmt.Println(people1) // [prakash sudeep udhay ram]

}
