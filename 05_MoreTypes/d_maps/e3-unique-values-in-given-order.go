package main

import (
	"fmt"
)

func main() {

	inputArr := []int{1, 1, 2, 3, 2, 0, 0}
	expectedOutputArr := []int{1, 2, 3, 0}
	fmt.Println(UniqueValuesInGivenOrder(inputArr), expectedOutputArr)
}

func UniqueValuesInGivenOrder(givenArr []int) []int {
	uniqueValues := make(map[int]int)
	var noDups []int
	for _, val := range givenArr {
		_, isKeyPresent := uniqueValues[val]
		if isKeyPresent == false {
			uniqueValues[val] = 1
			noDups = append(noDups, val)
		} else {
			uniqueValues[val]++
		}
	}

	return noDups
}
