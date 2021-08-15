package main

import "fmt"

/*
naked structs dont have type definition.
They are used only once.
Also, called inline struct
*/

func main() {
	firstStruct := struct {
		Name string
		Age  int
	}{
		Name: "udhay",
		Age:  18,
	} // populating with values

	fmt.Println("firstStruct=", firstStruct)
	fmt.Println("firstStruct.Name=", firstStruct.Name)

	firstStruct.Name = "prakash"
	fmt.Println("firstStruct=", firstStruct)

	secondStruct := struct {
		Name string
		Age  int
	}{}
	fmt.Println("secondStruct=", secondStruct)

	// empty or un-initialize. Golang will populate them automatically with default value
	// such as empty string or zero

	fmt.Printf("firstStruct struct : %+v\n", firstStruct)
	fmt.Printf("secondStruct struct : %+v\n", secondStruct)

}
