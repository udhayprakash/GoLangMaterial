package main

/*
Purpose: Higher-Order Function
	Passing Functions as an Argument to Another Function:
		- If a function is passed as an argument to another function,
		then such types of functions are known as a Higher-Order function.
		- This passing function as an argument is also known as a
		callback function or first-class function in the Go language.
*/
import (
	"fmt"
	"math"
)

// Sphere - Volume function takes
// a function as a argument
func Sphere(vol func(radius float64) float64) {
	fmt.Println("Volume of Sphere is:", vol(3))
}

// Main Function
func main() {

	volumeOfSphere := func(radius float64) float64 {
		result := 4 / 3 * math.Pi * radius * radius * radius
		return result
	}

	// Passing function as an
	// argumnt in Volume function
	Sphere(volumeOfSphere)
}
