package main

import "fmt"

type Point struct {
	X, Y float64
}

// Pointer Receivers
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
