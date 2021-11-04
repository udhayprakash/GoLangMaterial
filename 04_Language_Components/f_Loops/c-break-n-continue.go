package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n\nFull Loop")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	} // OUTPUT - 1 2 3 4 5 6 7 8 9 10

	fmt.Println("\n\nImportance of break")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // breaks out of for loop
		}
		fmt.Print(i, " ")
	} // OUTPUT - 1 2 3 4

	
	fmt.Println("\n\nImportance of continue")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // skips that loop
		}
		fmt.Print(i, " ")
	} // OUTPUT - 1 2 3 4 6 7 8 9 10


	fmt.Println("\n\nImportance of continue")
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			continue // skips the odd numbers
		}
		fmt.Print(i, " ")
	} // OUTPUT - 2 4 6 8 10

} 