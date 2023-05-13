package main

import (
	"fmt"
	_ "sort"
)

// sorting array in wave for which is
// arr[0] >= arr[1] <= arr[2] >= arr[3] <= arr[4] >= arr[5]
// and so on till array is sorted.

// Time Complexity: O(N)
// Auxiliary Space: O(1)
func sortInWaveForm(arr []int) {
	// Traverse all even elements
	for i := 0; i < len(arr)-1; i += 2 {
		// If current even element is smaller than previous
		if i > 0 && arr[i-1] > arr[i] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
		// If current even element is smaller than next
		if i < len(arr)-1 && arr[i] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}
func main() {
	// 5, 10, 1, 12, 45, 17, 89, 63
	// 10,23,54,7,4,9,53,12
	arr := []int{10, 23, 54, 7, 4, 9, 53, 12}
	fmt.Println("Before Tsunami Waves", arr)
	sortInWaveForm(arr)
	fmt.Println("After Tsunami Waves ", arr)
}

/*
OUTPUT:

$ go run 10_waveform_sorting.go
    Before Tsunami Waves [10 23 54 7 4 9 53 12]
    After Tsunami Waves  [23 10 54 4 9 7 53 12]
*/
