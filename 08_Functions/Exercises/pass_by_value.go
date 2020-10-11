package main

import "fmt"

// simple function to add 1 to a
func add1(a int) int {
	a = a + 1 // we change value of a
	return a  // return new value of a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // should print "x=3"

	x1 := add1(x) // call add1(x)

	fmt.Println("x+1 = ", x1) // should print "x+1 = 4"
	fmt.Println("x = ", x)    // should print "x=3"
}
