package main

import "fmt"

func main() {
	// To find the smallest number in the given list
	givenSlice := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var smallestVal int = givenSlice[0]
	for _, val := range givenSlice {
		if val < smallestVal {
			smallestVal = val
		}
	}
	fmt.Println("smallestVal:", smallestVal)
}
