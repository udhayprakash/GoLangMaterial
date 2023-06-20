package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("min(10, 20) = ", min(10, 20))
	fmt.Println("max(10, 20) = ", max(10, 20))
	fmt.Println()

	// Create two variables to store the minimum and maximum values.
	minVal := new(int)
	maxVal := new(int)

	// Call the `min` and `max` functions to get the minimum and maximum values.
	minPtr(minVal, 10, 20)
	maxPtr(maxVal, 10, 20)

	// Print the minimum and maximum values.
	fmt.Println("Minimum value:", *minVal)
	fmt.Println("Maximum value:", *maxVal)
}

func minPtr(minVal *int, a, b int) {
	// Check if the first value is less than the second value.
	if a < b {
		*minVal = a
	} else {
		*minVal = b
	}
}

func maxPtr(maxVal *int, a, b int) {
	// Check if the first value is greater than the second value.
	if a > b {
		*maxVal = a
	} else {
		*maxVal = b
	}
}
