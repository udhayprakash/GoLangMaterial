package main

import (
	"fmt"
	"reflect"
)

func main() {

	oddNumbers := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	//                     0  1  2  3  4   5   6   7

	// Method 1 - using for loop
	for index := 0; index < len(oddNumbers); index++ {
		// fmt.Println(index, oddNumbers[index])
		fmt.Printf("At position %d -> %2d\n", index, oddNumbers[index])
	}
	fmt.Println()

	// Method 2 - using range
	for range oddNumbers {
		fmt.Println("I am looping")
	}

	for index, value := range oddNumbers {
		fmt.Printf("At position %d -> %2d\n", index, value)
	}

	// ------- Generate even numbers between 1 & 100
	var evenNumbers [50]int // array
	fmt.Println("evenNumbers = ", evenNumbers, "type is ", reflect.TypeOf(evenNumbers).Kind())

	posCounter := 0
	for num := 1; num <= 100; num++ {
		if num%2 == 0 {
			evenNumbers[posCounter] = num
			posCounter++
		}
	}
	fmt.Println("evenNumbers=", evenNumbers)

}
