// Program to display the command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The command line args given are:")

	//fmt.Println("os.Args:", os.Args)
	// os.Args - It is a slice of strings
	programName := os.Args[0]
	arguments := os.Args[1:]

	fmt.Println("\tprogramName = ", programName)
	fmt.Println("\targuments   = ", arguments)

}

// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
