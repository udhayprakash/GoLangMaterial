package main

import "fmt"

var x string = "Global Scope"

func main() {
	var x string = "Block Scope"
	fmt.Println("In main()            ", x) // Local will be preferred
	my_function()
	another_function(x)
}

func my_function() {
	fmt.Println("In my_function()     ", x)
}
func another_function(x string) {
	fmt.Println("In another_function()", x)
}
