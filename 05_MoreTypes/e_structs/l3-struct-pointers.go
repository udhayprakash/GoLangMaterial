package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func createPoint(x, y int) *Point {
	p := Point{x, y}
	return &p
}
func main() {
	// var p1 = Point{1, 2}
	var p1 = new(Point)
	fmt.Println("p1 = ", p1) //  &{0 0}

	p1.x = 10
	fmt.Println("p1 = ", p1) //  &{10 0}

	(*p1).x = 5
	(*p1).y = 7
	fmt.Println("p1 = ", p1)          //  &{5 7}
	fmt.Println("(*p1).x =", (*p1).x) // 5
	fmt.Println("(*p1).y =", (*p1).y) // 7

	var point = Point{1, 2}
	var p2 = &point

	fmt.Println("\n\np2 =", p2)
	fmt.Println("(*p2).y =", (*p2).y)
	fmt.Println("(*p2).y =", (*p2).y)

	fmt.Println("createPoint(10, 20):", createPoint(10, 20)) //  &{10 20}

}
