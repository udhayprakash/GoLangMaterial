package main

// Methods with Pointer receivers as Functions

import (
	"fmt"
)

type PointA struct {
	X, Y float64
}

func (p *PointA) TranslateFunction(dx, dy float64) {
	// call by Reference
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := PointA{3, 4}
	fmt.Println("Point p = ", p) // {3 4}

	// TranslateFunc1(p, 7, 9)
	// TranslateFunc2(&p, 7, 9)

	p.TranslateFunction(7, 9)
	fmt.Println("After TranslateFunc1, p = ", p) //  {10 13

}
