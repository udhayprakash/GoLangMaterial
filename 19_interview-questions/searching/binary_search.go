package main

import "fmt"

func main() {
	numbers := []int{1, 4, 8, 12, 14, 16, 17, 21, 43, 56, 99}
	for _, val := range numbers {
		fmt.Println(
			search(val, numbers),
			searchV2(val, numbers),
			searchRecursive(val, numbers),
		)
	}
	fmt.Println(search(-1, numbers))
	fmt.Println(searchV2(15, numbers))
	fmt.Println(searchRecursive(101, numbers))

}

func search(want int, numbers []int) int {
	wantIndex := 0
	for len(numbers) > 1 {
		// Split the list in half by finding the middle index
		middle := len(numbers) / 2
		lowerHalf := numbers[0:middle]
		upperHalf := numbers[middle:]

		// Then we ask if our number is in the lower or upper half
		if want < numbers[middle] {
			numbers = lowerHalf
		} else {
			numbers = upperHalf
			// We need to account for the values in the lower half of the
			// list
			wantIndex += len(lowerHalf)
		}
	}
	if want == numbers[0] {
		return wantIndex
	}
	return -1
}

func searchV2(want int, numbers []int) int {
	lo, hi := 0, len(numbers)
	// The length of the range we are searching in must have more than one value,
	// otherwise we can stop searching and just look at that value.
	for (hi - lo) > 1 {
		mid := (lo + hi) / 2
		if want < numbers[mid] {
			// Use the lower half
			hi = mid
		} else {
			// Use the upper half
			lo = mid
		}
	}
	// We are done searching, and our value should be at index lo if it is in
	// the list.
	if lo >= len(numbers) || want != numbers[lo] {
		return -1
	}
	return lo
}

func searchRecursive(want int, numbers []int) int {
	// If our list has no numbers, we cannot find the value in the list.
	if len(numbers) == 0 {
		return -1
	}
	// If our list has 1 number, we can check to see if it is the number we want.
	if len(numbers) == 1 {
		if numbers[0] != want {
			return -1
		}
		return 0
	}
	// Otherwise we cut the list in half and recursively call search.
	mid := len(numbers) / 2
	var result int
	if want < numbers[mid] {
		result = searchRecursive(want, numbers[0:mid])
		if result < 0 {
			return -1
		}
		return result
	} else {
		result = searchRecursive(want, numbers[mid:])
		if result < 0 {
			return -1
		}
		return result + mid
	}
}