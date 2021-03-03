package main

import (
	"fmt"
	"image/color"
)

type Point2 struct {
	X, Y float64
}

type ColoredPoint2 struct {
	Point2 // anonymous fields - type of field is `named type` or `Point2er to named type`
	Color  color.RGBA
	isGood bool
}

func main() {
	var cp ColoredPoint2
	fmt.Println("cp =", cp) // cp = {{0 0} {0 0 0 0} false}
	fmt.Println()

	cp.X = 1
	fmt.Println("cp.X = ", cp.X)
	fmt.Println("cp.Y = ", cp.Y)
	fmt.Println("cp =", cp) // cp = {{0 0} {0 0 0 0} false}

	fmt.Println("cp.Y == cp.Point2.Y : ", cp.Y == cp.Point2.Y) // true
	fmt.Println()

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint2{Point2{1, 1}, red, true}
	var q = ColoredPoint2{Point2{5, 4}, blue, false}
	fmt.Println(p) // {{1 1} {255 0 0 255} true}
	fmt.Println(q) // {{5 4} {0 0 255 255} false}

}
