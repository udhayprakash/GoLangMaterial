package main

import (
	"fmt"
)

type Point struct {
	x, y int32
}

func main() {
	var p1 = new(Point)
	// p1.x and p1.y is 0
	(*p1).x = 5
	(*p1).y = 7

	fmt.Println("p1=", p1)
	fmt.Println("(*p1).x =", (*p1).x)
	fmt.Println("(*p1).y =", (*p1).y)

	var point = Point{1, 2}
	var p2 = &point

	fmt.Println("\n\np2 =", p2)
	fmt.Println("(*p2).y =", (*p2).y)
	fmt.Println("(*p2).y =", (*p2).y)
}
