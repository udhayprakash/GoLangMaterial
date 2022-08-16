package main

import (
	"fmt"
	. "sorting/insertionsort"
	. "sorting/quicksort"
)

func main() {
	arr := RandamArray(10)
	fmt.Println("Initial array     :", arr)
	fmt.Println("QuickSorted array :", QuickSort(arr))

	//arr = RandamArray(10)
	//fmt.Println("\nInitial array     :", arr)
	//fmt.Println("MergeSorted array :", MergeSort(arr))

	arr = RandamArray(10)
	fmt.Println("\nInitial array     :", arr)
	fmt.Println("InsertionSort array:", InsertionSort(arr))

}
