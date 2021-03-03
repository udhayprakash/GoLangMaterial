package main

import "fmt"

type myPoint struct {
	X, Y int
}

func main() {
	p1 := myPoint{1, 2}
	p2 := myPoint{2, 1}

	fmt.Println("p1 == p2 :", p1 == p2)
	fmt.Println("p1 != p2 :", p1 != p2)
	//fmt.Println("p1 > p2  :", p1 > p2) // invalid operation: p1 > p2 (operator > not defined on struct)
	//fmt.Println("p1 < p2  :", p1 < p2) // invalid operation: p1 < p2 (operator < not defined on struct)
}
