package main

import "fmt"

var x string = "Global Scope"

func main() {
	var x string = "Block Scope"

	// Local will be preferred
	fmt.Println("In main()       :", x, &x)

	myFunc()
	anotherFunc(x)
}

func myFunc() {
	fmt.Println("In myFunc()     :", x, &x)
}

func anotherFunc(x string) {
	// New Local variable will be created
	fmt.Println("In anotherFunc():", x, &x)
}
