package main

import (
	"fmt"
)

/*
Purpose: Binary Search
	- search strategy used to find elements within a list by consistently
      reducing the amount of data to be searched and thereby increasing the
      rate at which the search term is found.
	- To use a binary search algorithm, the list to be operated on must
      have already been sorted.
*/
func binarySearch(searchEle int, dataList []int) {
	lowerPos := 0
	higherPos := len(dataList) - 1
	for lowerPos <= higherPos {
		median := (lowerPos + higherPos) / 2
		if searchEle == dataList[median] {
			fmt.Printf("\t%3d is present at positon %2d\n", searchEle, median)
			return
		} else if searchEle < dataList[median] {
			higherPos = median - 1
		} else {
			lowerPos = median + 1
		}
	}
	fmt.Printf("\t%3d is NOT present in data: %2d\n", searchEle, dataList)
}

func main() {
	// Binary Search expects the data to be in sorted order
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println("items:", items)
	binarySearch(1, items)
	binarySearch(63, items)
	binarySearch(100, items)
	binarySearch(101, items)

	items = []int{1, 2, 9, 20}
	fmt.Println("items:", items)
	binarySearch(1, items)
	binarySearch(2, items)
	binarySearch(9, items)
	binarySearch(10, items)

}
