package main

import "fmt"

func main() {
	num1 := 20
	num2 := 12

	fmt.Printf("num1      =%v\n", num1)
	fmt.Printf("num1      =%b\n\n", num1)

	fmt.Printf("num2      =%v\n", num2)
	fmt.Printf("num2      =%b\n\n", num2)

	// Bitwise AND - &
	fmt.Println("num1 & num2 =", num1&num2)
	// Bitwise OR  - |
	fmt.Println("num1 | num2 =", num1|num2)
	// Bitwise XOR  - ^
	fmt.Println("num1 ^ num2 =", num1^num2)

}

/*
		128 64 32 16 8 4 2 1
	20    0  0  0  1 0 1 0 0
	12    0  0  0  0 1 1 0 0
	-------------------------
	&     0  0  0  0 0 1 0 0   => 4      bitwise AND
	|     0  0  0  1 1 1 0 0   => 28     bitwise OR
	^     0  0  0  1 1 0 0 0   => 24     bitwise XOR

*/

// Assignment - bitwise left shift(<<) and right shift(>>)