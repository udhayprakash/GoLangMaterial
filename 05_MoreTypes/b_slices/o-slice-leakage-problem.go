package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	printSlice(s)

	// Slice the slice to give it zero length.
	// Elements are not dropped
	s = s[:0]
	printSlice(s)

	// extend array to original size,
	// all elements are still there
	s = s[:5]
	printSlice(s)

	// Drop its first two values.
	// but why does it drop elements and we can't go back to the original?
	s = s[2:]
	printSlice(s)

	s = nil
	printSlice(s)

	// s = s[:3]
	// panic: runtime error: slice bounds out of range [:3] with capacity 0

	s = []int{9, 8, 7, 6, 5}
	printSlice(s) // NO More leakage

	s = s[:5]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
