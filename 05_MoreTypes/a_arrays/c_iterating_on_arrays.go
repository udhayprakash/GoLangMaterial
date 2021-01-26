package main

import "fmt"

func main() {

	oddNumbers := [...]int{1, 3, 5, 7, 9, 11, 13, 15}

	for index := 0; index < len(oddNumbers); index++ {
		//fmt.Println(" in loop", oddNumbers[index])
		fmt.Println("In Array ", oddNumbers, " at position", index, "value present is ", oddNumbers[index])
	}
	fmt.Println()

	// Iterating using range
	for index, value := range oddNumbers {
		fmt.Println("In Array ", oddNumbers, " at position", index, "value present is ", value)

	}
}
