package main

import (
	"fmt"
)

func selectionSort(randomSlice []int) {
	var size = len(randomSlice)

	for i := 0; i < size; i++ {
		var min = i

		// find the first, second, third, fourth...smallest value
		for j := i; j < size; j++ {
			if randomSlice[j] < randomSlice[min] {
				min = j
			}
		}

		// swap the smallest value the position of "i"
		randomSlice[i], randomSlice[min] = randomSlice[min], randomSlice[i]
	} // repeat till end of slice

}

func main() {
	var randomInt = []int{10, 9, 8, 22, 11, 23, 9}

	fmt.Println("Unsorted : ")
	fmt.Println(randomInt)

	fmt.Println("Selection Sorted : ")

	selectionSort(randomInt)
	fmt.Println(randomInt)

}
