package main

import "fmt"

/*
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted;
the variable will take the type of the initializer.
*/
func main() {
	// `var` declares 1 or more variables
	var a = "initial"
	fmt.Println("a=", a)

	var b = 12312323
	fmt.Println("b=", b)

	var c float64 = 123.45
	fmt.Println("c =", c)

	// More than One variable can be declared at once
	var d, e = 10, 20
	fmt.Println("d =", d, "e=", e)

	var f, g int = 30, 40
	fmt.Println("f =", f, "g =", g)

	// Go will infer the type of initialized variables.
	var h = true
	fmt.Println("h=", h)

	// Variables declared without initialization are _zero-valued_.
	// zero value for int -> 0
	// zero value for str -> ""
	var i int
	fmt.Println("i=", i)

	var (
		j float32
		k float64
	)
	fmt.Println("j:", j)
	fmt.Println("k:", k)

	// `:=` syntax for declaring and initializing a variable.
	// E.g. for `var l string = "apple"` in this case.
	l := "apple"
	fmt.Println("l:", l)

}
