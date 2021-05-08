package main

import "fmt"

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

func main() {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
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
