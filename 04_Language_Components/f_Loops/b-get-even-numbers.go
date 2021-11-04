package main

import (
	"fmt"
)

func main() {
	fmt.Println("Even number between 56 & 100:")

	for num := 56; num <= 100; num++ {
		if num % 2 == 0{
			fmt.Printf("%d,", num)
		}
	}
	fmt.Println()
} 