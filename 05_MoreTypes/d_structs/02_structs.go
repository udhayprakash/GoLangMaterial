package main
/*
Purpose: Structs
	- Go’s structs are typed collections of fields.
	- They’re useful for grouping data together to form
      records.
*/
import "fmt"

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
