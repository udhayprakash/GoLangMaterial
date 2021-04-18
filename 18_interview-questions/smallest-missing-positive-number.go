package main

import (
	"fmt"
	"sort"
)

/*
Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].


*/

func main() {
	arr1 := []int{1, 3, 6, 4, 1, 2}
	expectedOutput1 := 5
	result1 := Solution(arr1)
	fmt.Println(expectedOutput1, result1)

	arr2 := []int{1, 2, 3}
	expectedOutput2 := 4
	result2 := Solution(arr2)
	fmt.Println(expectedOutput2, result2)

	arr3 := []int{-1, -3}
	expectedOutput3 := 1
	result3 := Solution(arr3)
	fmt.Println(expectedOutput3, result3)

}

func filterPositive(A []int) []int {
	result := []int{}
	for _, val := range A {
		if val > 0 {
			result = append(result, val)
		}
	}
	return result
}

func Solution(A []int) int {
	A = filterPositive(A)
	if len(A) == 0 {
		return 1
	}
	sort.Ints(A)
	for index, val := range A {
		if index != 0 && val >= 0 && (val-A[index-1] > 1) {
			return A[index-1] + 1
		}

	}
	return A[len(A)-1] + 1
}
