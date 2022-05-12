package main

import "fmt"

// Summation, Average
func main() {
	var numbers [5]int
	fmt.Println("Before- numbers=", numbers)

	var summation int
	for i := range numbers {
		fmt.Scan(&numbers[i])
		summation += numbers[i]
	}
	fmt.Println("After - numbers=", numbers)
	fmt.Println("Sum of Values:", summation)

	fmt.Println("Avg of Values:", summation/len(numbers)) // type-casting loss
	fmt.Println("Avg of Values:", float32(summation)/float32(len(numbers)))

}

// Assignment - extend this to compute max, min and median and rank values
