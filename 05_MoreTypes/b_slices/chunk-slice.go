package main

import "fmt"

func main() {
	splitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)
	splitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)
	splitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	splitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4)
	splitArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
}

func splitArray(array []int, numberOfChunks int) {
	var result [][]int

	for i := 0; i < numberOfChunks; i++ {

		min := (i * len(array) / numberOfChunks)
		max := ((i + 1) * len(array)) / numberOfChunks

		result = append(result, array[min:max])
	}
	fmt.Println("result", result)
}
