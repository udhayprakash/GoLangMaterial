package main

import (
	"fmt"
	"math"
)

/*
Purpose: Interfaces
	- are named collections of method signatures.

*/

// structs
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
// methods to structs
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}


func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry){
	fmt.Println("g        :", g)
	fmt.Println("g.area() :", g.area())
	fmt.Println("g.perim():", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)

}

//Further ref: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

