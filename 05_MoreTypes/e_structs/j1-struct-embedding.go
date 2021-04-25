package main

import "fmt"

type Circle struct {
	X, Y, Radius int
}
type Wheel struct {
	X, Y, Radius, Spokes int
}

// struct embedding - place common features in another struct & using it
type Point struct {
	X, Y int
}
type Circle1 struct {
	Center Point
	Radius int
}
type Wheel1 struct {
	Circle Circle1
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	var w1 Wheel1
	w1.Circle.Center.X = 8
	w1.Circle.Center.Y = 8
	w1.Circle.Radius = 5
	w1.Spokes = 20
	fmt.Println(w1)
}
k-function-as-field.go