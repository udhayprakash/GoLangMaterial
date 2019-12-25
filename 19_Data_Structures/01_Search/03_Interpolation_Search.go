package main

import "fmt"
/*
Purpose: Interpolation Search
 - an improvement over Binary Search for instances,
    where the values in a sorted array are uniformly distributed.
 - Binary Search always goes to middle element to check.
 - interpolation search may go to different locations according
   the value of searchVal being searched.
*/
func interpolationSearch(dataList []int, searchVal int) int {
	lowerPos, higherPos := 0, len(dataList) -1
	min, max := dataList[0], dataList[higherPos]

	for {
		if searchVal < min {
			return lowerPos
		}
		if searchVal > max {
			return higherPos + 1
		}
	}
	// make a guess of the location
	var guess int
	if higherPos == lowerPos {
		guess = higherPos
	} else {
		size := higherPos - lowerPos
		offset := int(float64(size-1) * (float64(searchVal-min) / float64(max-min)))
		guess = lowerPos + offset
	}

	// maybe we found it?
	if dataList[guess] == searchVal {
		// scan backwards for start of value range
		for guess > 0 && dataList[guess-1] == searchVal {
			guess--
		}
		return guess
	}

	// if we guessed to high, guess lower or vice versa
	if dataList[guess] > searchVal {
		lowerPos = guess - 1
		max = dataList[higherPos]
	} else {
		lowerPos = guess + 1
		min = dataList[lowerPos]
	}
}

func main() {
	// Expects the elements to be in sorted order
	items := []int{-1, 0, 1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(items,1))
}
