package main

import "fmt"

func main() {
	// var number1 int
	// fmt.Println("Before, number1=", number1)

	// fmt.Print("Enter a number:")
	// fmt.Scanf("%d", &number1)
	// fmt.Println("After, number1=", number1)

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

$ go run p-runtime-input.go
Before, number1= 0
Enter a number:34
After, number1= 34

$ go run p-runtime-input.go
Before, number1= 0
Enter a number:34.43
After, number1= 34

$ go run p-runtime-input.go
Before, number1= 0
Enter a number:num234age56
After, number1= 0

$ go run p-runtime-input.go
Before, number1= 0
Enter a number:234age
After, number1= 234

CASE 2

$ go run p-runtime-input.go 
Enter no. of inches:12
inches: 12

$ go run p-runtime-input.go 
Enter no. of inches:12.233
inches: 12.233

$ go run p-runtime-input.go 
Enter no. of inches:-0.2323
inches: -0.2323


$ go run p-runtime-input.go 
Enter no. of inches:23123age
inches: 23123

$ go run p-runtime-input.go 
Enter no. of inches:213.23age
inches: 213.23

$ go run p-runtime-input.go 
Enter no. of inches:age213
inches: 0
*/
