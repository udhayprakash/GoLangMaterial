package main

import "fmt"

type Direction int

func main() {

	const (
		North Direction = iota
		East
		South
		West
	)

	fmt.Println("South =", South)
	fmt.Println("NorthEast =", North, East)

}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
