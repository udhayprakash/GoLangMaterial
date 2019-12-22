// Program to display the command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The command line args given are:")

	programName := os.Args[0]
	arguments := os.Args[1:]

	fmt.Println("\tprogramName = ", programName)
	fmt.Println("\targuments   = ", arguments)

}
