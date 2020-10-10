package main

import "fmt"

func main() {
	myArray := 	[...]int{1, 2, 3, 4}

	for index:= 0; index < len(myArray); index++ {
		// Accessing array elements
		fmt.Printf("a[%d] = %d\n", index, myArray[index])
	}
	fmt.Println()

	// Iterating using range
	for index, value := range myArray {
		// Accessing array elements
		fmt.Printf("a[%d] = %d\n", index, value)
	}
}
