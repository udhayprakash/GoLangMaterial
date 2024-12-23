package main

/*
Purpose: Methods
	- Go supports methods defined on struct types.
	- Go automatically handles conversion between values
      and pointers for method calls.
	- You may want to use a  pointer receiver type to avoid
      copying on method calls or to allow the method to
      mutate the receiving struct.
*/
import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area     :", r.area())
	fmt.Println("perimeter:", r.perimeter())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area     :", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
