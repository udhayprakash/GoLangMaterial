package main

import "fmt"

func main() {

	for count := 0; count < 10; count++ {
		fmt.Print(count, ", ")
	}

	fmt.Println()

	count := 0
	for count < 10 {
		fmt.Print(count, ", ")
		count++
	}

	fmt.Println()

	count = 0
	for {
		fmt.Print(count, ", ")
		if count >= 10 {
			break
		}
		count++
	}

	fmt.Println()

	// infinite condition
	// count = 0
	// for {
	// 	fmt.Print(count, ", ")
	// 	// if count >= 10 {
	// 	// 	break
	// 	// }
	// 	count++
	// }

	// for true {
	// 	fmt.Printf("This loop will run forever.\n")
	// }

	// Alternatively,
	for {
		fmt.Printf("This loop will run forever.\n")
	}
}
