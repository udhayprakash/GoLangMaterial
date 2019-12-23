package main
/*
Purpose: Methods
	- Methods are typically used to emulate class-like
      behavior on struct types. Methods do not have to
      receive a struct type.
	- Pointers can also be receiving types for methods.
	- Methods behave exactly the same as functions.

*/
import "fmt"

type Degrees struct {
	angle float64
}

//method
func (deg Degrees) toRad() float64 {
	return deg.angle * (3.14 / 180.0)
}

func main() {
	angle := Degrees { 90.0 }
	fmt.Println("Angle, in radians: ", angle.toRad())
}