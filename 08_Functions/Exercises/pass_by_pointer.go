package main

import "fmt"

// simple function to add 1 to a
func add2(a *int) int {
	*a = *a + 1 // we changed value of a
	return *a   // return new value of a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // should print "x=3"

	x1 := add2(&x) // call add1(&x) pass memory address of x

	fmt.Println("x+1 = ", x1) // should print "x+1 = 4"
	fmt.Println("x = ", x)    // should print "x=4"
}

/*
Advantages of pointers:
	• Allows us to use more functions to operate on one variable.
	• Low cost by passing memory addresses (8 bytes), copy is not an efficient way, both in terms
		of time and space, to pass variables.
	• string, slice and map are reference types, so they use pointers when passing to functions
		by default. (Attention: If you need to change the length of slice, you have to pass pointers
		explicitly)
*/
