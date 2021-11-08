package main

import "fmt"

func findFirstOccurrence(searchNum int, nums []int) {
	for index, value := range nums {
		if value == searchNum {
			fmt.Printf("%d found at index %d\n", value, index)
			return
		}
	}
	fmt.Printf("%d not found in %d\n", searchNum, nums)
}

func findAllOccurrences(searchNum int, nums []int) {
	foundIndices := []int{}
	for index, value := range nums {
		if searchNum == value {
			foundIndices = append(foundIndices, index)
		}
	}

	if foundIndices != nil {
		fmt.Printf("%d found at indices %d\n", searchNum, foundIndices)
	} else {
		fmt.Printf("%d not found in %d\n", searchNum, nums)
	}
}

func main() {
	findFirstOccurrence(45, []int{56, 67, 45, 90, 45})
	//                             0   1   2   3   4
	findFirstOccurrence(78, []int{78, 78, 78})
	findFirstOccurrence(87, []int{})
	fmt.Println()

	findAllOccurrences(45, []int{56, 67, 45, 90, 45})
	findAllOccurrences(78, []int{78, 78, 78})
	findAllOccurrences(87, []int{})

}
 