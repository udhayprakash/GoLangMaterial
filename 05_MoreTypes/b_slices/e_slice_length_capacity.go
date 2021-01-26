package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // Length   = 6  	capacity = 6  value    = [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // Length   = 0  	capacity = 6  value    = []

	// Extend its length.
	s = s[:4]
	printSlice(s) // Length   = 4  	capacity = 6  value    = [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // Length   = 2  	capacity = 4  value    = [5 7]

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%2d\t%v\n", i, cap(y), y)
		x = y
	}

}

func printSlice(mySlice []int) {
	fmt.Printf("\nLength   = %d  	capacity = %d  value    = %v", len(mySlice), cap(mySlice), mySlice)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
