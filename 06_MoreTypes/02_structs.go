package main

import "fmt"

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

}
