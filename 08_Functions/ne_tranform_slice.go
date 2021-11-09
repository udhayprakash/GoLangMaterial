package main

import "fmt"

func tranform(mySlice []int, f func(int) int) {
	for i := range mySlice {
		mySlice[i] = f(mySlice[i])
	}
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Initially, numbers=", numbers) // [1 2 3 4 5]

	// multiply each value with  3
	tranform(numbers, func(x int) int {
		return x * 3
	})
	fmt.Println("After *3 , numbers=", numbers) // [3 6 9 12 15]

	// Add with 10
	tranform(numbers, func(x int) int {
		return x + 10
	})
	fmt.Println("After +10, numbers=", numbers) // [13 16 19 22 25]

}
