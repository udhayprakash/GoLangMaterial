package main

import "fmt"

var x string = "Global Scope"

func main() {
	// var x string = "Block Scope"
	fmt.Println(x) // Locals will be preferred
}

func f() {
	fmt.Println(x)
}
