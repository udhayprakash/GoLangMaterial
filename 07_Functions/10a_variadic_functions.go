package main

import "fmt"

// Functions can have variadic parameters.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterate each value of the variadic.
	// The underbar here is ignoring the index argument of the array.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// Pass variadic value as a variadic parameter.
	//fmt.Println("params:", fmt.Println(myStrings)) // multiple-value fmt.Println() in single-value context
	fmt.Println("params:", fmt.Sprintln(myStrings))

	// ... - spread operator
	fmt.Println("params:", fmt.Sprintln(myStrings...))

}

func main() {
	learnVariadicParams()
	learnVariadicParams("Sun", "raises", "in", "the east")
}
