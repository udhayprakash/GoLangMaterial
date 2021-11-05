package main

import "fmt"

func main() {
	a := 111
	fmt.Println("Initally, a =", a) // 111

	defer fmt.Println("a = ", a) // 111 -- registers with current value

	a = 222
	fmt.Println("At end,   a =", a) // 222
}
 