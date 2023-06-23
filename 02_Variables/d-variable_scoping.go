package main

import "fmt"

var x string = "Global Scope"

func main() {
	fmt.Println("In main()       :", x, &x)

	var x string = "Block Scope"

	// Local will be preferred
	fmt.Println("In main()       :", x, &x)

	// function call
	funcWithNoInput()
	funcWithInput(x)

}

func funcWithNoInput() {
	fmt.Println("In funcWithNoInput():", x, &x)
}

func funcWithInput(x string) {
	// New Local variable will be created
	fmt.Println("In anotherFunc():", x, &x)
}
