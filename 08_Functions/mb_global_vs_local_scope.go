package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println("Inside function     : x=", x) // 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 0
	fmt.Println("Inside function     : x=", *xPtr) // 0
}

func main() {
	x := 5

	// case 1
	fmt.Println("\n\nCall By Value:")
	fmt.Println("Before function call: x=", x) // 5
	zero(x)
	fmt.Println("after  function call: x=", x) // 5


	// case 2
	fmt.Println("\n\nCall By Reference:")
	fmt.Println("Before function call: x=", x) // 5
	//zeroPointer(x) // cannot use x (type int) as type *int in argument to zeroPointer
	zeroPointer(&x)
	fmt.Println("after  function call: x=", x) // 0

}