package main

import "fmt"

/*
Purpose: Execution Order in switch Statements
	- First the switch expression is evaluated once.
	- Then case expressions are evaluated left-to-right and top-to-bottom:
		- the first one that equals the switch expression triggers execution
          of the statements of the associated case,
		- the other cases are skipped.
*/

// Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}


func main() {
	//Foo(1)
	//Foo(2131)
	//Foo(1)

	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}

}