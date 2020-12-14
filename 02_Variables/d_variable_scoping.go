package main

import "fmt"

var x string = "Global Scope"

func main() {
	var x string = "Block Scope"

	// Local will be preferred
	fmt.Println("In main()       :", x) 
	
	myFunc()
	anotherFunc(x)
}

func myFunc() {
	fmt.Println("In myFunc()     :", x)
}
func anotherFunc(x string) {
	fmt.Println("In anotherFunc():", x)
}
