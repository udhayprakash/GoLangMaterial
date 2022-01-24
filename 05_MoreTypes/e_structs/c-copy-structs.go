package main

import "fmt"

// type Point struct {
// 	X float64
// 	Y float64
// }

type Point struct {
	X, Y float64
}

// Structs are Value-types
func main() {
	p1 := Point{10, 20}

	// copying struct
	p2 := p1

	fmt.Println("p1 = ", p1, &p1.X) // {10 20} 0xc0000120c0
	fmt.Println("p2 = ", p2, &p2.X) //  {10 20} 0xc0000120d0

	p2.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p1 = ", p1) // {10 20}
	fmt.Println("p2 = ", p2) // {15 20}
}
