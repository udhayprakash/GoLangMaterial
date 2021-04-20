package main

import (
	"fmt"
)

func main() {
	fmt.Println("Even number between 56 & 100:")
	for i := 56; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
