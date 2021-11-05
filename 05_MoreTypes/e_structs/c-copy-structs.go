package main

// Structs are Value-types
import "fmt"

// type Point struct {
// 	X float64
// 	Y float64
// }

type Point struct {
	X, Y float64
}

func main() {
	// Structs are value types.
	p1 := Point{10, 20}

	// copying struct
	p2 := p1

	fmt.Println("p1 = ", p1, &p1.X) // {10 20} 0xc0000084f0
	fmt.Println("p2 = ", p2, &p2.X) //  {10 20} 0xc000008500

	p2.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p1 = ", p1) // {10 20}
	fmt.Println("p2 = ", p2) // {15 20}
}

// NOTE: Changes in one instance will not affect the other
 