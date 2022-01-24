package main

import "fmt"

/*
Purpose: Methods
	- Methods are typically used to emulate class-like
      behavior on struct types.
	- Methods do not have to receive a struct type.
	- Pointers can also be receiving types for methods.
	- Methods behave exactly the same as functions.

*/

type Degrees struct {
	angle     float64
	something string
}

// method
func (deg Degrees) toRadians() float64 {
	return deg.angle * (3.14 / 180.0)
}

func main() {
	angle1 := Degrees{90.0, "Inclination Angle"}
	fmt.Println("angle1 =", angle1) // {90 Inclination Angle}

	fmt.Println("angle1.toRadians() =", angle1.toRadians()) // 1.5699999999999998
}

// assignment: Round off the result to two digits, after decimal
