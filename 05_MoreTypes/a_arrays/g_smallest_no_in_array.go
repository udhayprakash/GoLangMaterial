package main

import (
	"fmt"
)

func main() {
	numbers := [7]int{11, 44, 22, -88, 33, 44, 56}
	fmt.Println("numbers=", numbers)

	// assuming first value is smallest
	minValue := numbers[0]
	for _, value := range numbers {
		if value < minValue {
			// updating smallest value
			minValue = value
		}
	}
	fmt.Printf("Smallest value in %v is %v\n", numbers, minValue)

}
