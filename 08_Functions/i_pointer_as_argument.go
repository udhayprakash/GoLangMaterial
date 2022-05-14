package main

import "fmt"

// Call by Value -- local changes won't reflect outside
func incrment(val int) int {
	val++
	return val
}

// Call by Reference -- local changes WILL reflect outside
func incrmentPointer(p *int) {
	// fmt.Println("p = ", p, "*p =", *p)
	*p++ // increments what p points to; does not change p
}

func main() {
	v := 1
	// Method 1 - updating the value
	v = incrment(v)
	fmt.Println("v =", v) // 2

	v = incrment(v)
	fmt.Println("v =", v) // 3

	// Method 2 - Using Pointer reference
	incrmentPointer(&v)
	fmt.Println("v =", v) // 4

	incrmentPointer(&v)
	fmt.Println("v =", v) // 5
}

/*
Advantages:
	- string, slice and map are reference types, so they use pointers
	when passing to functions by default.
	- Low cost by passing memory addresses (8 bytes), copy is not an
	efficient way, both in terms of time and space, to pass variables

	- Each time we take the address of a variable or copy a pointer,
  	we create new aliases or ways to identify the same variable.

*/
