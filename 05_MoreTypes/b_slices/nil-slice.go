package main

import "fmt"

func main() {
	var emptySlice []int
	fmt.Println("emptySlice       :", emptySlice)
	fmt.Println("emptySlice == nil:", emptySlice == nil)
	fmt.Println("len(emptySlice)  :", len(emptySlice))
	if emptySlice == nil {
		fmt.Println("Empty slice")
	}
}
