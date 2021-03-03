package main

import (
	"fmt"
	"reflect"
)

// Function Definition
//func helloWorld(){
//	fmt.Println("Hello")
//}

//func helloWorld(name string){
//	fmt.Println("Hello", name)
//}

//func helloWorld(name string) (string){
//	return "Hello " + name
//}

func helloWorld(name string) string {
	if name == "" {
		name = "GOLAND"
	}
	return "Hello " + name
}

func main() {
	// Function call
	fmt.Println("helloWorld(\"GoLANG\")=", helloWorld("GoLANG"))

	result := helloWorld("World")
	fmt.Println(reflect.TypeOf(result), result)

	result = helloWorld("")
	fmt.Println(reflect.TypeOf(result), result)

}
