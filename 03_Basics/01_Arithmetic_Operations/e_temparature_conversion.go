/*
Fahrenheit into Celsius
        C = (F - 32) * 5/9
*/
package main

import "fmt"

func main() {
	const F float32 = 100
	const C float32 = (F - 32) * 5 / 9

	fmt.Println("F = ", F)
	fmt.Println("C = ", C)

	var F1 float32 //= 100
	fmt.Println("Please Enter temperature in degrees Fahrenheit:")
	fmt.Scanf("%d", &F1)

	C1 := (F - 32) * 5 / 9

	fmt.Println("F1 = ", F1)
	fmt.Println("C1 = ", C1)
}

// Assignment : take two values in run time and compute +, -, *, /
