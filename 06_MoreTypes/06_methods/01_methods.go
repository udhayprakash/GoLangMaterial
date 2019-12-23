package main

import "fmt"

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

}
