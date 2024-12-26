package main

import "fmt"

/*
Go language supports FOUR types of for loop -
  - A complete, C-style for
  - A condition-only for  (equivalent to while loop)
  - An infinite for
  - for-range

C-style
========

	for initialization; condition; incre/decr {
		// zero or more statements
		}
*/
func main() {
	// Method 1
	var i int
	for i = 0; i <= 5; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Method 2
	// for var j int =0; j <= 5; j ++ {
	// 	var declaration not allowed in for initializer

	for j := 0; j <= 5; j++ {
		fmt.Printf("%d\t", j)
	}

	// Method 3 --> in while loop style
	k := 0
	for k <= 5 {
		fmt.Printf("%d\t", k)
		k++
	}
}
