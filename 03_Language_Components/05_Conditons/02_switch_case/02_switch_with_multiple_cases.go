package main

import (
	"fmt"
	"time"
)

func main() {
	// Here's a  `switch with Multiple cases`.
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 0:
		fmt.Println("It is Zero")
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even Number")
	}


	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() { // switch initializer statement
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
