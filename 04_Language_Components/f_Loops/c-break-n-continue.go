package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n\nFull Loop")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("\n\nImportance of break")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println("\n\nImportance of continue")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i, " ")
	}
}
