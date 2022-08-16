package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}
	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart
}

func main() {
	arr := RandomArray(10)
	fmt.Println("Initial Array:", arr)
	fmt.Println("sorted Array :", quickSort(arr))
}

// RandomArray - Generating a random array
func RandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
