package main

import "fmt"

func printASlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	printASlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

	// To specify a capacity, pass a third argument to make:
	b := make([]int, 0, 5)
	printASlice("b", b) // b len=0 cap=5 []

	c := b[:2]
	printASlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printASlice("d", d) // d len=3 cap=3 [0 0 0]
}
