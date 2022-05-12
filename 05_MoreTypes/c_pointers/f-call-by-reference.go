package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println("Before, x=", x, "&x=", &x) // 5 0xc0000100a8
	fmt.Println()

	zeroFunc(x)
	fmt.Println("After, x=", x, "&x=", &x) // 5
	fmt.Println()

	zeroPtrfunc(&x)
	fmt.Println("After, x=", x, "&x=", &x) // 0
}

// call by value -- changes in function, NOT reflected outside, func scope
func zeroFunc(z int) {
	fmt.Println("\tin func, before - z=", z, "&z=", &z) // 5
	z = 0
	fmt.Println("\tin func, After  - z=", z, "&z=", &z) // 0

}

// call by reference -- changes in function, reflected outside, func scope
func zeroPtrfunc(z *int) {
	fmt.Println("\tin func, before - z=", z, "*z=", *z) // 5
	*z = 0
	fmt.Println("\tin func, After  - z=", z, "*z=", *z) // 0

}
