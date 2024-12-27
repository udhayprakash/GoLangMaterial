package main

import (
	"fmt"
)

func main() {

	// comparision on arrays
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{2, 3, 1}
	fmt.Println("a1 == a2 :", a1 == a2) // false

	a3 := [4]int{1, 2, 3}
	fmt.Println("a3 = ", a3)
	// fmt.Println("a1 == a3 :", a1 == a3)
	// invalid operation: a1 == a3 (mismatched types [3]int and [4]int)

	a4 := [...]int{1, 2, 3}
	fmt.Println("a1 == a4 :", a1 == a4) // true
	fmt.Println()

	q := [3]int{1, 2, 3}
	fmt.Println("q =", q) // [1 2 3]

	// updating the array values, in full - Overwriting the values
	q = [3]int{44, 55, 66}
	fmt.Println("q =", q) // [44 55 66]

	// q = [2]int{777, 888}
	// cannot use [2]int{…} (value of type [2]int) as [3]int value in assignment

	q = [3]int{777, 888}
	fmt.Println("q =", q) // [777 888 0]

	// q = [5]int{11, 22, 33, 44, 55}
	// cannot use [5]int{…} (value of type [5]int) as [3]int value in assignment

	//q = [3]float64{4.4, 5.5, 6.6}
	// cannot use [3]float64 literal (type [3]float64) as type [3]int in assignment

	// Indexing
	fmt.Println("q =", q)          //  [777 888 0]
	fmt.Println("q[0] =", q[0])    // 777
	fmt.Println("q[1] =", q[1])    // 888
	fmt.Println("len(q)=", len(q)) // 3

	// updating a value in array  -- arrays are MUTABLE
	q[1] = 89
	fmt.Println("q[1] =", q[1]) // 89
	fmt.Println("q =", q)       // 	[777 89 0]

	// Slicing -- will get till the last position, excluding last position
	fmt.Println("q[0:2] =", q[0:2]) // [777 89]

	// updating a slice of array
	// q[0:2] = [2]int{11, 22}
	// cannot assign to q[0:2] (neither addressable nor a map index expression)
}

/*
NOTE:
1. arrays are mutable
2. we can update individual indexed elements, but not a slice of them
3. Overwriting/updating is possible with same dimension; values can be less than or equal to the dimension
*/
