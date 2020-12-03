package main

import "fmt"

func main() {

	integersToSort := make([]int, 0, 1)
	// prompt user for 10 integers (use a loop)
	for count := 0; count < 10; count++ {
		fmt.Printf("Please enter integer %d: ", count + 1)
		integersToSort = AddIntegerToSlice(integersToSort)
	}

	fmt.Printf("Original slice: %v", integersToSort)
	BubbleSort(&integersToSort)
	fmt.Printf("\nSorted slice: %v", integersToSort)
}

func AddIntegerToSlice(integersToSort []int) []int {

	var userInput int

	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Printf("An error occurred when trying to scan user input:\n\t%s", err)
	} else {
		integersToSort = append(integersToSort, userInput)
	}
	return integersToSort

}

func BubbleSort(integersToSort *[]int) {

	length := len(*integersToSort)

	for i := 0; i < length - 1; i++ {
		for j := 0; j < length - i - 1; j ++ {
			if (*integersToSort)[j] > (*integersToSort)[j+1] {
				Swap(integersToSort, j)
			}
		}
	}
}

func Swap(integersToSort *[]int, index int) {
	temp := (*integersToSort)[index]
	(*integersToSort)[index] = (*integersToSort)[index+1]
	(*integersToSort)[index+1] = temp
}
