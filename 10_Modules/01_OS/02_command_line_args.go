// Program to display the command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The command line args given are:")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("argument ", i, "is", os.Args[i])
	}

}
