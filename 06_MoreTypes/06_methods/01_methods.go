package main

import (
	"fmt"
	"math"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

// Method with receiver `Point`
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{2.0, 4.0}

	fmt.Println("p =", p)
	fmt.Println("Is Point p located above line y = 1.0 ? : ", IsAboveFunc(p, 1))
	fmt.Println("Is Point p located above line y = 1.0 ? : ", p.IsAbove(1))

	p2 := Point{3.4, 5.6}
	fmt.Println("Distance(p, p2):", Distance(p, p2)) // function call
	fmt.Println("p.Distance(p2) :", p.Distance(p2))  // method call
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	// Here, parameter p is called method's receiver
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// In Go, we don't use special variable like this or self
// for the receivers.

/*
The expression p.Distance is calle d a selector, because it selects the appropriate Distance
method for the receiver p of type Point. Selectors are also used to select fields of struct types,
as in p.X. Since methods and fields in habit the same namespace, declaring a method X on the
struct type Point would be ambiguous and the compiler will reject it.
*/
