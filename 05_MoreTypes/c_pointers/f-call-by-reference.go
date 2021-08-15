package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println("Before, x=", x, "&x=", &x)

	zero(&x)
	fmt.Println("After, x=", x, "&x=", &x)
}

func zero(z *int) {
	fmt.Println("in func, before - z=", z, "*z=", *z)
	*z = 0
	fmt.Println("in func, After - z=", z, "*z=", *z)
}
// cALL by reference - changes within function are reflected outside
//                     the function scope.