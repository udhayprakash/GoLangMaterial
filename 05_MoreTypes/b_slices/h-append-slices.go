package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s) // [  ]

	s[0], s[1] = "a", "b"
	s[2] = "c"
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
