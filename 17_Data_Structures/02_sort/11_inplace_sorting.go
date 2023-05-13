package main

import (
	"fmt"
)

// inplace sorting, within a single for loop
func sortInts(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func sortStrings(arr []string) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	// non-reversed array
	arr1 := []int{1, 2, 45, 65, 32, 34}
	// output: [1 2 32 34 45 65]
	fmt.Printf("non reversed array: %v\n", arr1)
	// calling sortInts function
	sortInts(arr1)
	fmt.Printf("sorted array: %v\n", arr1)

	// non-reversed array
	arr2 := []string{"1", "2", "45", "65", "32", "34"}
	// output: [1 2 32 34 45 65]
	fmt.Printf("non reversed array: %v\n", arr2)
	// calling sortStrings function
	sortStrings(arr2)
	fmt.Printf("sorted array: %v\n", arr2)
}
