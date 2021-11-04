package main

import "fmt"

func main() {

	oddNumbers := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	//                     0  1  2  3  4   5   6   7

	// Method 1 - using for loop
	for index := 0; index < len(oddNumbers); index++ {
		fmt.Printf("At position %d -> %2d\n", index, oddNumbers[index])
	}
	fmt.Println()

	// Method 2 - using range
	for index, value := range oddNumbers {
		fmt.Printf("At position %d -> %2d\n", index, value)
	}

	// ------- Generate even numbers between 1 & 100
	var evenNumbers1 [50]int
	fmt.Println("evenNumbers1=", evenNumbers1)

	evenNumbers2 := make([]int, 50)
	fmt.Println("evenNumbers2=", evenNumbers2)

	count := 0
	for num := 1; num <= 100; num++ {
		if num%2 == 0 {
			evenNumbers2[count] = num
			count++
		}
	}
	fmt.Println("evenNumbers2=", evenNumbers2)

}
 