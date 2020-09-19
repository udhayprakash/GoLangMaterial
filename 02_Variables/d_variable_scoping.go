package main

import "fmt"

var x string = "Global Scope"

func main() {
	var x string = "Block Scope"
	fmt.Println("In main()           ", x) // Local will be preferred
	myFunction()
	anotherFunction(x)
}

func myFunction() {
	fmt.Println("In myFunction()     ", x)
}
func anotherFunction(x string) {
	fmt.Println("In anotherFunction()", x)
}
