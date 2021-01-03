package main

import (
	"fmt"
)

func main() {
	myFirstStruct := struct {
		Name string
		Age  int
	}{
		Name: "Caleb",
		Age:  102,
	} // populate with value

	mySecondStruct := struct {
		Name string
		Age  int
	}{}

	// empty or un-initialize. Golang will populate them automatically with default value
	// such as empty string or zero

	fmt.Printf("myFirstStruct struct : %+v\n", myFirstStruct)
	fmt.Printf("mySecondStruct struct : %+v\n", mySecondStruct)

}
