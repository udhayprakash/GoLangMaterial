package main

import "fmt"

/*
Fahrenheit into Celsius
        C = (F - 32) * 5/9
*/

func main() {
	const F float32 = 100
	const C float32 = (F - 32) * 5 / 9

	fmt.Printf(`
		F = %f
		C = %f
	`, F, C)

	var F1 float32
	fmt.Print("Please Enter temperature in degrees Fahrenheit:")
	fmt.Scanf("%f", &F1)

	C1 := (F1 - 32) * 5 / 9
	fmt.Printf(`
		F1 = %f 
		C1 = %f`, F1, C1)
}

/*
asignments:
	1) WAP to implement celcius to fahrenheit conversion
	2) WAP to take two values in runtime and compute +, -, *, /

*/
