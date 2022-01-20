package main

import "fmt"

// data type alias
type Direction int

func main() {
	const (
		north int = iota
		east
		south
		west
	)
	fmt.Println(north, east, south, west)

	const (
		North Direction = iota
		East
		South
		West
	)

	fmt.Println(North, East, South, West)

}
