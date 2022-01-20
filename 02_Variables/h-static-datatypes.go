package main

import "fmt"

var data float32 = 6.7

// data = 3453 // here, we can only declare; not update values

func main() {

	fmt.Printf("data = %v type = %T\n", data, data) // package-level value

	var data string
	fmt.Printf("data = %v type = %T\n", data, data)

	data = "first"
	fmt.Printf("data = %v type = %T\n", data, data)

	data = "juayguya gsd8768763*&^*&^(*)("
	fmt.Printf("data = %v type = %T\n", data, data)
	// NOTE: Modified value should be of same data type

	data = "a"
	fmt.Printf("data = %v type = %T\n", data, data)

	// data = 'a'
	// cannot use 'a' (type untyped rune) as type string in assignment

	// data = 3453
	// cannot use 3453 (type untyped int) as type string in assignment

	// var data int = 3453

}
    //NOTE: A variable cant be declared more than once
