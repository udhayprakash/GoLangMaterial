package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	fmt.Println("Before- numbers=", numbers)

	var sum int
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
		sum += numbers[i]
	}
	fmt.Println("After - numbers=", numbers)
	fmt.Println("Sum of Values:", sum)

	fmt.Println("Avg of Values:", sum/len(numbers)) // type-casting loss
	fmt.Println("Avg of Values:", float32(sum)/float32(len(numbers)))

}

// Assignment - extend this to compute max, min and median  and rank values
