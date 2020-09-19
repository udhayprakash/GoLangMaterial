package main

import "fmt"

// data type aliases
type Direction int

func main() {
	const (
		North Direction = iota
		East
		South
		West
	)

	fmt.Println("South     =", South)
	fmt.Println("NorthEast =", North, East)

}
