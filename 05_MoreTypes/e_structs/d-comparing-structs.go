package main

import (
	"fmt"
	"reflect"
)

type myPoint struct {
	X, Y int
}

type AnotherPoint struct {
	A, B int
}

func main() {
	p1 := myPoint{1, 2}
	p2 := myPoint{2, 1}

	fmt.Println("p1 == p2 :", p1 == p2)
	fmt.Println("p1 != p2 :", p1 != p2)
	// fmt.Println("p1 > p2  :", p1 > p2) // invalid operation: p1 > p2 (operator > not defined on struct)
	// fmt.Println("p1 < p2  :", p1 < p2) // invalid operation: p1 < p2 (operator < not defined on struct)

	p3 := myPoint{1, 2}
	fmt.Println("p1 == p3 :", p1 == p3)

	// Equality Check
	// Two structs are equal if all their corresponding fields are equal.
	p4 := AnotherPoint{1, 2}
	// fmt.Println("p1 == p4 :", p1 == p4)
	// mismatched types myPoint and AnotherPoint
	fmt.Println(reflect.TypeOf(p1), reflect.TypeOf(p1).Kind())
	fmt.Println(reflect.TypeOf(p4), reflect.TypeOf(p4).Kind())

}
