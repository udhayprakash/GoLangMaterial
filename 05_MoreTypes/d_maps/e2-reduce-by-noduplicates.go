package main

// Write a function that accepts a slice of integers and returns a slice of integers
// containing all the integers from the input slice that do not duplicate.
// In the event there are no duplicates, or the input is an empty slice and empty slice should be returned.
// You may safely assume the function will never receive a nil slice.

import (
	"fmt"
)

func GetNoDups(givenSlice []int) []int {
	if len(givenSlice) == 0 {
		return []int{}
	}
	uniqueNums := make(map[int]int)
	maxCount := 0
	for _, num := range givenSlice {
		uniqueNums[num]++
		if uniqueNums[num] > maxCount {
			maxCount = uniqueNums[num]
		}
	}
	if maxCount == 1 {
		//  No value duplicate
		return []int{}
	}
	// Else get all non-duplicate nums
	var nondups []int
	for num, count := range uniqueNums {
		if count == 1 {
			nondups = append(nondups, num)
		}
	}
	return nondups
}

func main() {

	inputSlice := []int{0, 2, 6, 2, 1, 4, 5, 3, 4, 9, 11, 2, 3, 9, 42, 22, 9, 0, 1, 3, 99}
	expectedSlice := []int{5, 6, 11, 22, 42, 99}

	fmt.Println(GetNoDups(inputSlice), expectedSlice)
	fmt.Println(GetNoDups([]int{}), []int{})     // empty
	fmt.Println(GetNoDups([]int{1, 2}), []int{}) // empty
}
