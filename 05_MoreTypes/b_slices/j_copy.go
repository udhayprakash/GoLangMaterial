package main

import "fmt"

func main() {
	// Array doesnt support copy, whereas slices will

	// Single-dimension slices
	myArray1 := []int{11, 22, 33, 44}
	myArray2 := make([]int, 4)

	// Assignment will lead to leakage
	//myArray2 = myArray1

	// Deep Copy
	copy(myArray2, myArray1) // destSlice, sourceSlice

	fmt.Println("Before :", myArray1, myArray2)

	myArray1[2] = 20
	fmt.Println("After  :", myArray1, myArray2)
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
