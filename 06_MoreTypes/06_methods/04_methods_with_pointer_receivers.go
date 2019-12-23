package main

import (
	"fmt"
)

// With a value receiver, the method operates on a copy of
// the value passed to it. Therefore, any modifications done
// to the receiver inside the method is not visible to the caller.

// Methods with pointer receivers can modify the value to which
// the receiver points. Such modifications are visible to the
// caller of the method as well -

type Point struct {
	X, Y float64
}

/*
  Translates the current Point, at location (X,Y), by dx along the x axis and dy along the y axis
  so that it now represents the point (X+dx,Y+dy).
*/
func (p *Point) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := Point{3, 4}
	fmt.Println("Point p = ", p)

	p.Translate(7, 9)
	fmt.Println("After Translate, p = ", p)
}
