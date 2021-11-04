package main

import "fmt"

func main() {
	// var numbers [10]int
	numbers := make([]int, 10)

	// WRONG APPROACH
	// fmt.Println("Enter 10 numbers:")
	// fmt.Scanf("%v", &numbers)
	// fmt.Println("numbers=", numbers)


	// Correct way 
	for index := 0; index < len(numbers); index++ {
		fmt.Printf("\nEnter %d position value:", index)
		fmt.Scanf("%d", &numbers[index])

	}
	fmt.Println("numbers=", numbers)

}
// Asignment - try the same using range 