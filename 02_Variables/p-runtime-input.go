package main

import "fmt"

func main() {
	var number1 int
	fmt.Println("Before, number1=", number1)

	fmt.Print("Enter a number:")
	fmt.Scanf("%d", &number1)
	fmt.Println("After, number1=", number1)

	// -------------------
	var inches float32

	fmt.Print("Enter no. of inches:")
	fmt.Scanf("%f", &inches)
	fmt.Println("inches:", inches)

}
/*
NOTE: If inappropriate date type value is given,
		it will return zero value only
	  If you give overflowing value, it will retain the zero 
	    value itself

OUTPUT:

	Before, number1= 0
	Enter a number:765
	After, number1= 765
	Enter no. of inches:inches: 0

	~go run p-runtime-input.go        
	Before, number1= 0
	Enter a number:65765.76
	After, number1= 65765
	Enter no. of inches:inches: 76

	~go run p-runtime-input.go
	Before, number1= 0
	Enter a number:5465.767.788
	After, number1= 5465
	Enter no. of inches:inches: 767.788

	~go run p-runtime-input.go
	Before, number1= 0
	Enter a number:jgjhgasd87ashggjas23.23
	After, number1= 0
	Enter no. of inches:inches: 0

	~go run p-runtime-input.go
	Before, number1= 0
	Enter a number:234 234.34 45.56
	After, number1= 234
	Enter no. of inches:inches: 234.34
*/