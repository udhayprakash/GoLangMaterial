package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		// Modifying the value variable will not modify
		// the value in the compound type.
		v *= 2
	}
	fmt.Println(evenVals) // [2 4 6 8 10 12]
}
