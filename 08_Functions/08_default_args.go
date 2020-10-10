package main
/*
Purpose: Default arguments for functions
    are not supported in golang.
*/
import (
	"fmt"
	"reflect"
)

func hello(name string) (string) {
	// set default values
	if name == "" {
		name = "GoLang!"
	}

	return "Hello " + name
}

func main() {
	result := hello("World")
	fmt.Println(reflect.TypeOf(result), result)

	result = hello("")
	fmt.Println(reflect.TypeOf(result), result)
}
