package main

import (
	"fmt"
)

func main() {
	// completed
	fmt.Println(computation([]int{1, 7})) // [1, 8]
	fmt.Println(computation([]int{1, 9})) // [2, 0]
	fmt.Println(computation([]int{1, 0, 9})) // [1, 1, 0]
	// working
	fmt.Println(computation([]int{9, 9})) // [1, 0, 0]
}

func computation(given []int) []int {
	if given[len(given) -1] != 9 {
		given[len(given) -1]++
		return given
	}
	for i:= len(given)-1; i >=0; i-- {
		if given[i] == 9{
			given[i] = 0
			if given[i-1] == 9 {
				// given = append(1, 0, given[i:len(i)-1])
			}   
			return given
		}
		
	}
	return given
}
