// Insertion Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
The idea of swapping adjacent elements to sort a list of items can also be used to implement the insertion sort. In the insertion sort algorithm, we assume that a certain portion of the list has already been sorted, while the other portion remains unsorted. With this assumption, we move through the unsorted portion of the list, picking one element at a time. With this element, we go through the sorted portion of the list and insert it in the right order so that the sorted portion of the list remains sorted.
*/
func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionsort(slice)
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

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}