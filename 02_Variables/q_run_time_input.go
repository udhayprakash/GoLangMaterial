package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")

	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println("You entered :", input)

	// ------------------
	fmt.Println("Enter no. of inches:")
	var inches float32
	fmt.Scanf("%f", &inches)

	feet := inches / 2.5
	fmt.Println("inches:", inches, "feet:", feet)

}
