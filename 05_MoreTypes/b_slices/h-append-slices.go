package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s) // [  ]

	s[0] = "a"
	s[1], s[2] = "b", "c"
	fmt.Println("set:", s)    // [a b c]
	fmt.Println("get:", s[2]) // c

	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)      // [a b c d e f]
	fmt.Println("len:", len(s)) // 6

	// Dynamic slice
	c := make([]string, len(s))
	copy(c, s)             // safe copy
	fmt.Println("cpy:", c) // [a b c d e f]
	fmt.Println()

	// In this case, changes in one slice, were not leaked into another slice
	c[3] = "D"
	fmt.Println("s =", s) // [a b c d e f]
	fmt.Println("c =", c) // [a b c D e f]

	// Declaring 2-D slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [[0] [1 2] [2 3 4]]

}
