package main

import "fmt"

func main(){
	// `var` declares 1 or more variables
	var a = "initial"
	fmt.Println("a=", a)

	var b = 12312323
	fmt.Println("b=", b)

	var c float32 = 12312.123
	fmt.Println("c=", c)

	// More than One variable can be declared at once
	var d, e = 1, 2
	fmt.Println("d=", d, "e=", e)

	var f, g int = 3, 4
	fmt.Println("f=", f, "g=", g)

	// Go will infer the type of initialized variables.
	var h = true
	fmt.Println("h=", h)

	// Variables declared without initialization are _zero-valued_.
	// zero value for int -> 0
	// zero value for str -> ""
	var i int
	fmt.Println("i:", i)

	// `:=` syntax for declaring and initializing a variable.
	// E.g. for `var f string = "apple"` in this case.
	j := "apple"
	fmt.Println("j:", j)

	var (
		k float32
		l float64
	)
	fmt.Println("k:", k)
	fmt.Println("l:", l)

}
// NOTE: In Go, variables are explicitly declared and used by
// the compiler to e.g. check type-correctness of function calls.

