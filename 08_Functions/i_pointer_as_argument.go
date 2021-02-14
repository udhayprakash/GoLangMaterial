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

// Each time we take the address of a variable or copy a pointer,
// we create new aliases or ways to identify the same variable.
