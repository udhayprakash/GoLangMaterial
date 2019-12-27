package main

import "fmt"
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	array1 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("initially     :", array1) // [5 4 3 2 1 0]
	reverse(array1[:])
	fmt.Println("After Reversal:", array1) // [5 4 3 2 1 0]

	// Rotate slice1 left by two positions.
	slice1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("Initially            :", slice1) // [0 1 2 3 4 5]
	reverse(slice1[:2])
	fmt.Println("After slice1[:2]     :", slice1) // [0 1 2 3 4 5]
	reverse(slice1[2:])
	fmt.Println("After slice1[2:]     :", slice1) // [0 1 2 3 4 5]
	reverse(slice1)
	fmt.Println("After reverse(slice1):", slice1) // [2 3 4 5 0 1]
}
