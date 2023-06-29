package main

import "fmt"

func main() {
	var userInput float32

	fmt.Printf("Please enter a floating point number: ")
	// fmt.Scanf("%f", &userInput)

	val, err := fmt.Scanf("%f", &userInput)
	fmt.Println("val", val)  // 1 - when correct, 0 when error
	fmt.Println("err", err)  // nil - when correct,  err when error
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("userInput = ", userInput)
	}
}

/*
$ go run q-untime-input.go 
Please enter a floating point number: 23.4
val 1
err <nil>
userInput =  23.4

$ go run q-untime-input.go 
Please enter a floating point number: sdas
val 0
err strconv.ParseFloat: parsing "": invalid syntax
strconv.ParseFloat: parsing "": invalid syntax
*/