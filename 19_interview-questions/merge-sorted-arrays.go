package main

/*
Given two sorted lists, 
write a code to merge them into one single sorted list.
{1, 4, 7,8,6} {3, 4, 5,9,2}

*/
import "fmt"

func merge(nums1 []int, nums2 []int) []int {
	merged := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			if len(merged) == 0 || merged[len(merged)-1] != nums1[i] {
				merged = append(merged, nums1[i])
			}
			i++
		} else if nums1[i] > nums2[j] {
			if len(merged) == 0 || merged[len(merged)-1] != nums2[j] {
				merged = append(merged, nums2[j])
			}
			j++
		} else {
			if len(merged) == 0 || merged[len(merged)-1] != nums1[i] {
				merged = append(merged, nums1[i])
			}
			i++
			j++
		}
	}

	for i < len(nums1) {
		if len(merged) == 0 || merged[len(merged)-1] != nums1[i] {
			merged = append(merged, nums1[i])
		}
		i++
	}

	for j < len(nums2) {
		if len(merged) == 0 || merged[len(merged)-1] != nums2[j] {
			merged = append(merged, nums2[j])
		}
		j++
	}
	return merged
}

func main() {
	n1 := []int{1, 4, 7, 8, 6}
	n2 := []int{2, 3, 4, 5, 9}
	fmt.Println("merged = ", merge(n1, n2))

	n1 = []int{1, 1, 3, 3, 4}
	n2 = []int{2, 2, 3, 4, 5}
	fmt.Println("merged = ", merge(n1, n2))

	n1 = []int{1, 1, 3}
	n2 = []int{2, 2, 3, 4, 5}
	fmt.Println("merged = ", merge(n1, n2))

	n1 = []int{1, 1, 3, 3, 4}
	n2 = []int{2, 2, 3}
	fmt.Println("merged = ", merge(n1, n2))
}
