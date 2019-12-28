package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	myList := []string{"1", "10", "11", "2", "3", "4", "5", "6", "7", "8", "9"}

	fmt.Printf("Before: %v\n", myList)

	// Pass in our list and a func to compare values
	sort.Slice(myList, func(i, j int) bool {
		numA, _ := strconv.Atoi(myList[i])
		numB, _ := strconv.Atoi(myList[j])
		return numA < numB
	})

	fmt.Printf("After : %v\n", myList)
}
