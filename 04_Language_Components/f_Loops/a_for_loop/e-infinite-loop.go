package main

import (
	"fmt"
)

func main() {

	count := 0
	for {
		count++
		fmt.Println(count, "This loop will run forever.")
		if count == 10 {
			break
		}
	}

	// for true {
	// 	fmt.Printf("This loop will run forever.\n")
	// }

	// Alternatively,
	for {
		fmt.Printf("This loop will run forever.\n")
	}
}
