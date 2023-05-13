package main

import (
	"fmt"
)

// // non optimized bubbleSort algorithm
// func bubbleSort(arr []int) {
//     for i := 0; i < len(arr)-1; i++ {
//         for j := i + 1; j < len(arr); j++ {
//             if arr[i] > arr[j] {
//                 arr[i], arr[j] = arr[j], arr[i]
//             }
//         }
//         fmt.Println(arr)
//     }
// }

// optimized bubbleSort algorithm
func bubbleSort3(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swapped = true
			}
		}
		fmt.Println(arr)
		if swapped == false {
			break
		}
	}
}

func main() {
	arr := []int{13, 2, 5, 14, 10, 1}
	fmt.Println("before bubble sort:", arr)
	bubbleSort3(arr)
	fmt.Println("after bubble sort:", arr)
}
