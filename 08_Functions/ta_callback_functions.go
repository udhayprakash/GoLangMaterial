/*
Purpose: Higher Order Functions
	Returning Functions From Another Functions
		- If a function returns another function, then such
		types of functions are known as a Higher-Order function.
		- It is also known as the First-Class function.
*/

// Golang program to illustrate how to pass
// a function returns another function
package main

import (
	"fmt"
	"math"
)

// Sphere -  Volume function returns
// an anonymous function
func Sphere() func(radius float64) float64 {

	result := func(radius float64) float64 {
		volume := 4 / 3 * math.Pi * radius * radius * radius
		return volume
	}

	return result

}

// Main Function
func main() {

	sVol := Sphere()
	fmt.Println("Volume of Sphere is:", sVol(5))
}
