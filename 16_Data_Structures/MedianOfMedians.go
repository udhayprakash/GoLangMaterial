package main

import (
	"fmt"
	"sort"
)
// In computer science, the median of medians is an approximate (median) selection algorithm, frequently used to supply a good pivot for an exact selection algorithm, mainly the quick-select, that selects kth smallest element. Here is source code of the Go Program to Median of Medians to find the Kth Smallest element.
func medianOfMedians(sliceList []int, k, r int) int {

	num := len(sliceList)
	if num < 10 {
		sort.Ints(sliceList)
		return sliceList[k-1]
	}
	med := (num + r - 1) / r

	medians := make([]int, med)

	for i := 0; i < med; i++ {
		v := (i * r) + r
		var arr []int
		if v >= num {
			arr = make([]int, len(sliceList[(i*r):]))
			copy(arr, sliceList[(i*r):])
		} else {
			arr = make([]int, r)
			copy(arr, sliceList[(i*r):v])
		}
		sort.Ints(arr)
		medians[i] = arr[len(arr)/2]
	}
	pivot := medianOfMedians(medians, (len(medians)+1)/2, r)

	var leftSide, rightSide []int

	for i := range sliceList {
		if sliceList[i] < pivot {
			leftSide = append(leftSide, sliceList[i])
		} else if sliceList[i] > pivot {
			rightSide = append(rightSide, sliceList[i])
		}
	}

	switch {
	case k == (len(leftSide) + 1):
		return pivot
	case k <= len(leftSide):
		return medianOfMedians(leftSide, k, r)
	default:
		return medianOfMedians(rightSide, k-len(leftSide)-1, r)
	}
}

func main() {
	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}
	sort.Ints(intSlice)

	for _, j := range []int{5, 10, 15, 20} {
		i := medianOfMedians(intSlice, j, 5)
		fmt.Println(j, "smallest element = ", i)
		v := intSlice[j-1]
		fmt.Println("arr[", j-1, "] = ", v)
		if i != v {
			fmt.Println("Oops! Algorithm is wrong")
		}
	}
}