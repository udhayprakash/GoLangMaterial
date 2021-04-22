package main

import "fmt"

/*
Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes
*/

func main() {
	names := [4]string{
		"Udhay",
		"Rob",
		"Prakash",
		"Thompson",
	}
	fmt.Println(names) // [Udhay Rob Prakash Thompson]
	//                       0    1    2       3
	//                           xxx

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [Udhay Rob] [Rob Prakash]

	// slices are mutable- but changes will be on array underlying
	b[0] = "XXX"
	fmt.Println(a, b)  // [Udhay XXX] [XXX Prakash]
	fmt.Println(names) // [Udhay XXX Prakash Thompson]

}
