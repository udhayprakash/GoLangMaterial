// Selection Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
This sorting algorithm begins by finding the smallest element in an array and interchanging it with data at, for instance, array index [0]. Starting at index 0, we search for the smallest item in the list that exists between index 1 and the index of the last element. When this element has been found, it is exchanged with the data found at index 0. We simply repeat this process until the list becomes sorted.
*/
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}