/*
You are given two arbitrarily large numbers,
stored one digit at a time in a slice.
The first must be added to the second,
and the second must be reversed before addition.

The goal is to calculate the sum of those two sets of values.

IMPORTANT NOTE:
- The input can be any lengths (i.e: it can be 20+ digits long).
- num1 and num2 can be different lengths.

Sample Inputs:
num1 = 123456
num2 = 123456

Sample Output:
Result: 777777

We would also like to see a demonstration of appropriate unit tests
for this functionality.
*/

package main

import (
	"fmt"
)

func Add(num1 []int, num2 []int) string {
	// Reversing the num2 slice.
	// As inplace changes not possible for slice, created new rev slice for clean code
	num2_rev := []int{}
	for i, _ := range num2 {
		num2_rev = append(num2_rev, num2[len(num2)-1-i])
	}

	resultStr := ""
	if len(num1) <= len(num2_rev) {
		for i := 0; i < len(num1); i++ {
			resultStr += fmt.Sprintf("%d", num1[i]+num2_rev[i])
		}
		for _, v1 := range num2_rev[len(num1):] {
			resultStr += fmt.Sprintf("%d", v1)
		}
	} else {
		for i := 0; i < len(num2_rev); i++ {
			resultStr += fmt.Sprintf("%d", num1[i]+num2_rev[i])
		}
		for _, v1 := range num1[len(num2_rev):] {
			resultStr += fmt.Sprintf("%d", v1)
		}
	}
	return resultStr
}

func main() {
	num1 := []int{}
	num2 := []int{}

	num1length := 6
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	result := Add(num1, num2)

	fmt.Println("Result:", result)
	fmt.Println()

	fmt.Println("Equal Sizes case        :", Add([]int{1, 2, 3}, []int{1, 2, 3}) == "444")
	fmt.Println("num1 len < num2 len case:", Add([]int{1, 2, 3}, []int{1, 2, 3, 4}) == "5551")
	fmt.Println("num1 len > num2 len case:", Add([]int{1, 2, 3, 4}, []int{1, 2, 3}) == "4444")
	fmt.Println("Equal Sizes case        :", Add([]int{1, 1, 2}, []int{7, 3, 8}) == "949")
	fmt.Println("one val each case       :", Add([]int{1}, []int{3}) == "4")
}

/*
1 2 3
3 2 1  // 123
-----
4 4 4


1 2 3 4
3 2 1  // 123
-----
4 4 4 4

1 2 3
4 3 2 1  // 1234
-----
5 5 5 1

123456
654321 //123456
===============
777777
*/
