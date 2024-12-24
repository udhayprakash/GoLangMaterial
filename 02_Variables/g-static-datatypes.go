package main

import "fmt"

var data float32 = 6.7

// data = 34.53 // here, we can only declare; not update values

func main() {

	fmt.Printf("data = %v type = %T\n", data, data) // package-level value

	data = 34.53
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
	//  cannot use 'a' (untyped rune constant 97) as string value in assignment

	// data = 3453.234
	// cannot use 3453.234 (untyped float constant 3453.23) as string value in assignment

	// var data int = 3453
	// data redeclared in this block
}

//NOTE: A  variable cant be declared more than once, within a SCOPE
