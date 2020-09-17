package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 6
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// -------------------
	var v1 int = 42
	var v2 float64 = float64(v1)
	var v3 uint = uint(v2)
	fmt.Println(v1, v2, v3)

	// ----------------------
	p1 := 42
	p2 := float64(p1)
	p3 := uint(p2)
	fmt.Println(p1, p2, p3)
	// Unlike in C, in Go assignment between items of different
	// type requires an explicit conversion.
}
