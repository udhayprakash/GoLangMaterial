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

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{3, 4}
	fmt.Println("Point            p = ", p)

	p.Translate(7, 9)
	fmt.Println("After Translate, p = ", p)

	//=========================================
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println("*r=", *r) // "{2, 4}"

	r1 := Point{1, 2}
	pptr := &r1
	pptr.ScaleBy(2)
	fmt.Println("r1=", r1) // "{2, 4}"

	r2 := Point{1, 2}
	(&r2).ScaleBy(2)
	fmt.Println("r2=", r2) // "{2, 4}"
}
