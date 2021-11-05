package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println("Before, x=", x, "&x=", &x) // 5
	fmt.Println()

	zeroFunc(x)
	fmt.Println("After, x=", x, "&x=", &x) // 5
	fmt.Println()

	zeroPtrfunc(&x)
	fmt.Println("After, x=", x, "&x=", &x) // 0

}

// call by value -- changes in function, NOT reflected outside, func scope
func zeroFunc(z int) {
	fmt.Println("in func, before - z=", z, "&z=", &z)
	z = 0
	fmt.Println("in func, After  - z=", z, "&z=", &z)

}

// call by reference -- changes in function, reflected outside,  func scope
func zeroPtrfunc(z *int) {
	fmt.Println("in func, before - z=", z, "*z=", *z)
	*z = 0
	fmt.Println("in func, After  - z=", z, "*z=", *z)

}
 