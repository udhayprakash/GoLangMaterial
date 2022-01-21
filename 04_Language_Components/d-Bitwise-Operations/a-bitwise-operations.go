package main

import "fmt"

/*
Bitwise Operations

    &   binary AND    -> If all are 1, result is 1
    |   binary OR     -> If all are 0, only then result is 0
    ^   binary XOR    -> If total count of 1's are odd, it results in 1, else 0
    ~   binary 1's complement
    <<  binary Left Shift
    >>  binary right shift



*/
func main() {
	num1, num2 := 20, 12

	fmt.Printf("num1      =%v\n", num1)   // 10 -- decimal
	fmt.Printf("num1      =%b\n\n", num1) // 10100 -- binary

	fmt.Printf("num2      =%v\n", num2)   // 12
	fmt.Printf("num2      =%b\n\n", num2) // 1100

	// Bitwise AND - &
	fmt.Println("num1 & num2 =", num1&num2) // 4
	// Bitwise OR  - |
	fmt.Println("num1 | num2 =", num1|num2) // 28
	// Bitwise XOR  - ^
	fmt.Println("num1 ^ num2 =", num1^num2) // 24
}

/*
		128 64 32 16 8 4 2 1
	20    0  0  0  1 0 1 0 0
	12    0  0  0  0 1 1 0 0
	-------------------------
	&     0  0  0  0 0 1 0 0 => 4  bitwise AND
	|     0  0  0  1 1 1 0 0 => 28 bitwise OR
	^     0  0  0  1 1 0 0 0 => 24 bitwise XOR


            128    64    32    16    8   4   2   1
 10                                  1   0   1   0
 20                             1    0   1   0          1st shift
 40                       1     0    1   0              2nd shift
 80                 1     0     1    0                  3rd shift
 160        1       0     1     0                       4th shift


NOTE:
A left shift by n bits is equivalent to multiplication by pow(2, n) without overflow check.
A right shift by n bits is equivalent to division by pow(2, n) without overflow check.

*/
// Assignment - bitwise left shift(<<) and right shift(>>)
