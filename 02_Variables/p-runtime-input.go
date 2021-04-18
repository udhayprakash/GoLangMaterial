package main

import "fmt"

func main() {
	var number1 int

	fmt.Println("Before, number1=", number1)
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &number1)
	fmt.Println("After, number1=", number1)

	// ------------------
	fmt.Println("Enter no. of inches:")
	var inches float32
	fmt.Scanf("%f", &inches)
	fmt.Println("inches:", inches)

	// fmt.Scanln()
}

// NOTE: If inappropriate date type value is given,
// it will return zero value only
// if you give overflowing value, it will retain the zero value itself
/*
OUTPUT:
	Before, number1= 0
	786
	After, number1= 786

	~go run p-runtime-input.go
	Before, number1= 0
	45.67
	After, number1= 45

	~go run p-runtime-input.go
	Before, number1= 0
	jhg
	After, number1= 0

*/
