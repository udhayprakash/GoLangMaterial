/*
CODILITY:
Write a function Solution that, given an integer N, returns the smallest
number with the same number of digits. You can assume N is between 1 and 10^9 (billion).

For example, given N = 125, the function should return 100
			 given N = 10, the function should return 10
			 Given N = 1, the function should return 0
*/

package main

import (
	"fmt"
	"math"
)

func SmallestPositiveNumber(N int) int {
	if N <= 9 {
		return 0
	}
	digits := 0
	for N > 0 {
		N = N / 10
		digits++
	}
	return int(math.Pow(10, float64(digits-1)))
}

func main() {
	for input, expectedOutput := range map[int]int{
		4751: 1000,
		3891: 1000,
		189:  100,
		125:  100,
		100:  100,
		37:   10,
		10:   10,
		1:    0,
	} {
		fmt.Println(input, SmallestPositiveNumber(input) == expectedOutput, expectedOutput)
	}
}
