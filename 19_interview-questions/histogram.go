package main

import (
	"fmt"
)

const SIZE = 10

var n = [SIZE]int{20, 7, 11, 13, 1, 3, 30, 4, 9, 2}

func main() {
	fmt.Printf("%sss\n", "Element", "Value", "Histogram")

	for i := 0; i <= SIZE-1; i++ {
		fmt.Printf("%7dd ", i, n[i])

		for j := 1; j <= n[i]; j++ {
			//fmt.Printf("%c", '*')
			fmt.Printf("%c", '∎')
		}
		fmt.Println()
	}
}
