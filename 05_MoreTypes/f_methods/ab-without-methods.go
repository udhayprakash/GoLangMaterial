package main

// Methods with Pointer receivers as Functions

import (
	"fmt"
)

type PointA struct {
	X, Y float64
}

// Pointer receivers as function argument
func TranslateFunc1(p PointA, dx, dy float64) {
	// call by value
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func TranslateFunc2(p *PointA, dx, dy float64) {
	// call by reference
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := PointA{3, 4}
	fmt.Println("Point p = ", p) // {3 4}

	TranslateFunc1(p, 7, 9)
	fmt.Println("After TranslateFunc1, p = ", p) // {3 4}

	TranslateFunc2(&p, 7, 9)
	fmt.Println("After TranslateFunc1, p = ", p) //  {10 13}
}
