package main
/*
range iterates over elements in a variety of data structures.
*/

import "fmt"

func main() {
	// Looping over an array
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for loopIndex, num := range nums {
		fmt.Println("At loopIndex", loopIndex, ": value", num, "is found")
	}

	// Looping over a map - both key and value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Looping over a map - only key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Looping over a string
	for i, c := range "go" {
		fmt.Println(i, c)
		// first value is the starting byte index of the rune
		// second the rune itself.
	}
}