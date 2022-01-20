package main

import "fmt"

func main() {
	// variable declaration & initialization

	// Method 1
	var x string
	x = "Hello World"
	fmt.Println("x =", x)

	// Method 2
	var y string = "Hello World"
	fmt.Println("y =", y)

	// Method 3- type is not necessary because the Go compiler is able to
	//    infer the type based on the literal value assigned to the variable.
	z := "Hello World"
	// Type inference - variable's type is inferred from value on right hand side.
	fmt.Println("z =", z)

	i := 42           // int
	f := 3.142        // float64 - an IEEE-754 64-bit floating point number.
	g := 0.867 + 0.5i // complex128 - represented internally with two float64's.
	fmt.Println(i, f, g)

	// Bulk variable declaration cum initialization
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)

	
	myRune := 'a'
	fmt.Println(myRune) // 97

	myRune = 'A'
	fmt.Println(myRune) // 65

}
