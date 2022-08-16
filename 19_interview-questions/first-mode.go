package main

import "fmt"

func main() {
	testValues := [][]int{
		[]int{5, 5, 2, 2, 1},
		[]int{5, 5, 2, 2, 2, 1},
		[]int{5, 2, 1},
		[]int{-1, -1, -1},
		[]int{-1},
		[]int{99},
	}
	for _, values := range testValues {
		fmt.Printf("GetFirstMode(%2v)=%v\n", values, GetFirstMode(values))
	}

}

func GetFirstMode(vals []int) int {
	freqs := make(map[int]int)
	for _, val := range vals {
		_, isKeyPresent := freqs[val]
		if isKeyPresent == true {
			freqs[val] += 1
		} else {
			freqs[val] = 1
		}
	}
	for key, count := range freqs {
		if count > 1 {
			for _, val := range vals {
				if key == val {
					return val
				}
			}
		}
	}
	return -1 // no mode (no duplicate value)
}
