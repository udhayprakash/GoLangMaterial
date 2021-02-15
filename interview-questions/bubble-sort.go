package main

import "fmt"

func bubbleSort(values []int) []int {
	swapped := true
	for swapped == true {
		swapped = false
		for i := 0; i < len(values); i++ {
			if i != 0 && values[i-1] > values[i] {
				values[i-1], values[i] = values[i], values[i-1]
				swapped = true
			}
		}
	}
	return values
}

func main() {
	unorderedArray := []int{12, 45, -3, 5, 1, 6}
	fmt.Println("unorderedArray=", unorderedArray)

	orderedArray := bubbleSort(unorderedArray)
	fmt.Println("orderedArray  =", orderedArray)
}
