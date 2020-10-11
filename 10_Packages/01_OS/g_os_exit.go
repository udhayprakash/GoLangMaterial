package main

/*
Purpose: os.Exit
	- Use os.Exit to immediately exit with a given status.
	- Note that unlike e.g. C, Go does not use an integer
      return value from main to indicate exit status.
    - If you’d like to exit with a non-zero status you
      should use os.Exit.
*/
import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("hello - start")
	os.Exit(1)
	// unreachable code
	fmt.Println("hello - end")
}

func main() {
	fmt.Println("main - start")
	defer fmt.Println("This is deferred statement")

	hello()
	//os.Exit(7687) // program execution terminates

	// Below code is not reachable
	fmt.Println("main - end")

}
