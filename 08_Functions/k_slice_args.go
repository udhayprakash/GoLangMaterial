package main

import "fmt"

func PassArrayInt10(ar [10]int) {
	for _, val := range ar {
		fmt.Print(val, ", ")
	}
	fmt.Println()
}

func PassArrayInt3(ar [3]int) {
	for _, val := range ar {
		fmt.Print(val, ", ")
	}
	fmt.Println()
}

func PassSliceInt(sl []int) {
	for _, val := range sl {
		fmt.Print(val, ", ")
	}
	fmt.Println()
}

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
	foundIndices := []int{} // slice
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
	PassArrayInt10([10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 00})
	PassArrayInt10([10]int{11, 22, 33})

	// PassArrayInt10([3]int{11, 22, 33})
	// cannot use [3]int{…} (value of type [3]int) as type [10]int in argument to PassArrayInt10

	PassArrayInt3([3]int{11, 22, 33})

	fmt.Println()
	PassSliceInt([]int{1, 2, 3})
	PassSliceInt([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 00})
	PassSliceInt([]int{})

	ar1 := [...]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 00}
	PassSliceInt(ar1[:])

	//=========================
	findFirstOccurrence(45, []int{56, 67, 45, 90, 45})
	//                             0   1   2   3   4
	findFirstOccurrence(78, []int{78, 78, 78})
	findFirstOccurrence(87, []int{})
	fmt.Println()

	findAllOccurrences(45, []int{56, 67, 45, 90, 45})
	findAllOccurrences(78, []int{78, 78, 78})
	findAllOccurrences(87, []int{})

}


// assignment - enhance this function findAllOccurrences to return the indices
