package main

import "fmt"

func main() {
	// Array doesnt support copy, whereas slices will

	// Single-dimension slices
	slice1 := []int{11, 22, 33, 44}
	slice2 := make([]int, 4)

	// Assignment will lead to leakage
	//slice2 = slice1

	// Deep Copy
	copy(slice2, slice1) // destSlice, sourceSlice

	fmt.Println("Before :", slice1, slice2)

	slice1[2] = 20
	fmt.Println("After  :", slice1, slice2)
	fmt.Println()

	// ================================================
	// two-dimensional Slices
	sourceSlice := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	DestinationSlice := make([][]int, 3)
	// Deep Copy
	copy(DestinationSlice, sourceSlice)
	fmt.Println("sourceSlice     :", sourceSlice)
	fmt.Println("DestinationSlice:", DestinationSlice)

	fmt.Println("sourceSlice[0]   =", sourceSlice[0])
	fmt.Println("sourceSlice[0][0]=", sourceSlice[0][0])

	sourceSlice[0][0] = 99
	fmt.Println("sourceSlice     :", sourceSlice)
	fmt.Println("DestinationSlice:", DestinationSlice)

}
