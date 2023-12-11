/*
How to concatenate two matrices horizontally
*/
package main

import (
	"fmt"
)

func main() {
	// Sample matrices
	matrix1 := [][]int{{1, 2}, {3, 4}}
	matrix2 := [][]int{{5, 6}, {7, 8}}

	// Function to concatenate matrices horizontally
	result := concatenateMatricesHorizontally(matrix1, matrix2)

	// Display the concatenated matrix
	for _, row := range result {
		fmt.Println(row)
	}
}

// Function to concatenate matrices horizontally
func concatenateMatricesHorizontally(matrix1, matrix2 [][]int) [][]int {
	if len(matrix1) != len(matrix2) {
		panic("Matrices should have the same number of rows")
	}

	result := make([][]int, len(matrix1))
	for i := range result {
		// Create a new row with the combined length
		result[i] = make([]int, len(matrix1[i])+len(matrix2[i]))

		// Copy elements from the first matrix
		copy(result[i][:len(matrix1[i])], matrix1[i])

		// Copy elements from the second matrix
		copy(result[i][len(matrix1[i]):], matrix2[i])
	}
	return result
}
