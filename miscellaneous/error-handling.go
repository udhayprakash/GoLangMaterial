package main

import (
	"fmt"
)

func rec() {
	if a := recover(); a != nil {
		fmt.Println(a)
	}
}

func Foo(arg0 *string, arg1 *string) {
	defer rec()

	if arg0 == nil {
		panic("Error")
	}
	if arg1 == nil {
		panic("Error:nil")
	}
	fmt.Println("Arg0 %s Arg1 %s", *arg0, *arg1)
	fmt.Println("Pass")
}

func main() {
	valS := "SomeText"
	Foo(&valS, nil)
	fmt.Println("Complete")
}
