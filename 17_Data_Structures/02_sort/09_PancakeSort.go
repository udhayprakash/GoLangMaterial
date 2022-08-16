// Pancake Sort in Golang
package main

import "fmt"

// Pancake sorting is the colloquial term for the mathematical problem of sorting a disordered stack of pancakes in order of size when a spatula can be inserted at any point in the stack and used to flip all pancakes above it. A pancake number is the maximum number of flips required for a given number of pancakes. Here is source code of the Go Program to implement Pancake Sorting Algorithm.
func main() {
	list := data{28, 11, 59, -26, 503, 158, 997, 193, -23, 44}
	fmt.Println("\n--- Unsorted --- \n\n", list)

	list.pancakesort()
	fmt.Println("\n--- Sorted ---\n\n", list, "\n")
}

type data []int32

func (dataList data) pancakesort() {
	for uns := len(dataList) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, dataList[0]
		for i := 1; i <= uns; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}
		// move to final position in two flips
		dataList.flip(lx)
		dataList.flip(uns)
	}
}

func (dataList data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}
