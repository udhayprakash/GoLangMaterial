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
			break // breaks out of the current for loop
		}

		fmt.Print(i, " ")
	} //  OUTPUT - 1 2 3 4

	fmt.Println("\n\nImportance of continue")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // skips that current loop
		}
		fmt.Print(i, " ")
	} // OUTPUT - 1 2 3 4 6 7 8 9 10

	fmt.Println("\n\nImportance of continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skips the even numbers
		}
		fmt.Print(i, " ")
	} // OUTPUT - 1 3 5 7 9

	// fizzbuzz problem
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
