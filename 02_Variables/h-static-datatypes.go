package main 

import "fmt"

var data float32 = 6.7

func main(){

	fmt.Printf("data = %v type = %T\n", data, data) // package-level value

	var data string 
	fmt.Printf("data = %v type = %T\n", data, data)

	data = "first"
	fmt.Printf("data = %v type = %T\n", data, data)

	data = "juayguyagsd8768763*&^*&^(*)("
	fmt.Printf("data = %v type = %T\n", data, data)
	// NOTE: Modified value should be of same data type

	data = "a"
	fmt.Printf("data = %v type = %T\n", data, data)

	// data = 'a'
	// fmt.Printf("data = %v type = %T\n", data, data)
	// error: incompatible types in assignment 


	// data = 3453
	// fmt.Printf("data = %v type = %T\n", data, data)
	// error: incompatible types in assignment (cannot use type int as type string)

	// var data int = 3453
	// fmt.Printf("data = %v type = %T\n", data, data)
	// error: redefinition of âdataâ
}

	//NOTE: A variable cant be declared more than once
 