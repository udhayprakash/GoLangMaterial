package main

/*
Purpose: Structs
	- Go’s structs are typed collections of fields.
	- They’re useful for grouping data together to form
      records.
*/
import "fmt"

type VertexCoordinate struct {
	X int
	Y int
}

func main() {
	fmt.Println("VertexCoordinate{1, 2} =", VertexCoordinate{1, 2})

	// Struct fields are accessed using a dot.
	v := VertexCoordinate{1, 2}
	fmt.Println("v.X = ", v.X)
	fmt.Println("v.Y = ", v.Y)

}