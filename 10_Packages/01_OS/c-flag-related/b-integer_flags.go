package main

import (
	"flag"
	"fmt"
)

func main() {
	// Step 1: defining the flags
	var num1 int
	flag.IntVar(&num1, "num1", 999, "Number 1") // stores in variable

	number2 := flag.Int("num2", 55, "Number 2 value") // creates pointer to variable

	// Step 2: command line flags parsing
	flag.Parse()

	// Step 3: Using the flag values
	fmt.Printf("type = %4T num1    = %#[1]v \n", num1)
	// fmt.Println()
	// fmt.Printf("type = %4T number2 = %#[1]v \n", number2)
	fmt.Printf("type = %4T number2 = %#[1]v \n", *number2)

}

/*
OUTPUT:
=======
	~go run b-integer_flags.go
	type =  int num1    = 999
	type = *int number2 = (*int)(0xc0000120c8)
	type =  int number2 = 55

	~go run b-integer_flags.go -num1=222
	type =  int num1    = 222
	type =  int number2 = 55

	~go run b-integer_flags.go -num1=222 -num2=333
	type =  int num1    = 222
	type =  int number2 = 333

	~go run b-integer_flags.go -h
	Usage of C:\...\exe\b-integer_flags.exe:
	-num1 int
			Number 1 (default 999)
	-num2 int
			Number 2 value (default 55)

*/
