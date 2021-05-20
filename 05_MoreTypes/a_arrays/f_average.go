package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	fmt.Println("Before- numbers=", numbers)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println("After - numbers=", numbers)

	// Summation & Average
	var sum int
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of Values:", sum)
	fmt.Println("Avg of Values:", sum/len(numbers))
	fmt.Println("Avg of Values:", float32(sum)/float32(len(numbers)))
}
