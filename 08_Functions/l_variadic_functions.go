package main

import "fmt"

// Functions can have variadic parameters.
//func funcWithAnyNoOfArgs(){
//	fmt.Println("funcWithAnyNoOfArgs - start")
//}

//func funcWithAnyNoOfArgs(name string){
//	fmt.Println("funcWithAnyNoOfArgs - start", name)
//}

//func funcWithAnyNoOfArgs(name interface{}){
//	fmt.Println("funcWithAnyNoOfArgs - start", name)
//}

// ... - spread operator
func funcWithAnyNoOfArgs(name ...interface{}) {
	fmt.Println("funcWithAnyNoOfArgs - start")

	fmt.Printf("value :%v\n", name)
	fmt.Printf("type  :%T\n\n", name)

	for val := range name {
		fmt.Println("val:", val)
	}
}

func main() {
	// Function call
	funcWithAnyNoOfArgs("Udhay")
	funcWithAnyNoOfArgs(2342342)
	funcWithAnyNoOfArgs("Udhay", "Prakash")
	funcWithAnyNoOfArgs("Udhay", "Prakash", 123123, 213.23, []int{1, 2, 3})
}

/*
Define a function sum(), which can add any number of value,
either int or float or together
if any other data type is given it should throw an error

HINT: func sum(nums ...int) {}

*/
