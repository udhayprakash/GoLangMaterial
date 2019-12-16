/*
Fahrenheit into Celsius 
        C = (F - 32) * 5/9
*/
package main

import "fmt"

func main() {
	const F float32 = 100
	const C float32 = (F - 32) * 5/9

	fmt.Println("F = ", F)
	fmt.Println("C = ", C)

}

