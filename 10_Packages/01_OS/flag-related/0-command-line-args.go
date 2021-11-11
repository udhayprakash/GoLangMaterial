// Program to display the command line arguments
package main

import (
	"fmt"
	"os"
	"reflect"
)

// Program to display the command line arguments
func main() {
	// var num1 int
	// fmt.Print("Enter something:")
	// fmt.Scanf("%d", &num1)

	fmt.Println("The command line args given are:")

	// fmt.Println("os.Args:", os.Args)

	fmt.Println("No of args:", len(os.Args))
	fmt.Println("this file :", os.Args[0])
	fmt.Println("Arguments:", os.Args[1:])

	fmt.Println(reflect.TypeOf(os.Args))  // []string
}

// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.