package main

import "fmt"

// It is perfectly safe for a function to return
// the address of a local variable.
func f() *int {
	v := 1
	return &v
}

func main() {
	var p = f()
	fmt.Println("f():", p, *p)


	p = f()
	fmt.Println("f():", p, *p)

	p = f()
	fmt.Println("f():", p, *p)

	fmt.Println(f() == f()) // "false"
}