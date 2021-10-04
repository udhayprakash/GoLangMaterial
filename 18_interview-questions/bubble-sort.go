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

func bubbleSort2(array[] int)[]int {
	for i:=0; i< len(array)-1; i++ {
	   for j:=0; j < len(array)-i-1; j++ {
		  if (array[j] > array[j+1]) {
			 array[j], array[j+1] = array[j+1], array[j]
		  }
	   }
	}
	return array
 }
func main() {
	unorderedArray := []int{12, 45, -3, 5, 1, 6}
	fmt.Println("unorderedArray=", unorderedArray)

	orderedArray := bubbleSort(unorderedArray)
	fmt.Println("orderedArray  =", orderedArray)

	orderedArray2 := bubbleSort2(unorderedArray)
	fmt.Println("orderedArray2 =", orderedArray2)
}
