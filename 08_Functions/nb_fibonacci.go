package main

import "fmt"

// fibonacci Series: 0 1 1 2 3 5 8 13 .....
func fibonacci() func() int {
	a, b := 0, 1
	return func() int { // anonymous functions
		// a = b
		// b = a + b
		a, b = b, a+b // unpacking
		return a
	}
}

func main() {
	f := fibonacci()

	fmt.Println("f() =", f())
	fmt.Println("f() =", f())
	fmt.Println("f() =", f())

	for i := 0; i < 10; i++ {
		fmt.Println("f():", f())
	}
}
