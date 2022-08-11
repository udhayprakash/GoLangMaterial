package main

/*
Purpose: os.Exit
	- Use os.Exit to immediately exit with a given status.
	- Note that unlike e.g. C, Go does not use an integer
      return value from main to indicate exit status.
    - If you’d like to exit with a non-zero status you
	  should use os.Exit.
	- This is a well established method in POSIX systems,
	  whereby a program can return a 0-255 integer to indicate
	  if the program ran successfully, and if not, why not.
	- Common codes
	  --------------------------------------------------
		code 			message
	  --------------------------------------------------
		 1				General error
		 2				misuse of shell builtins
		 127			Command not found
		 128 			Fatal error signal
		 130 			Ctrl+C termination
*/
import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("\thello - start")
	defer fmt.Println("\t\thello - defer")
	// os.Exit(1)
	// unreachable code
	fmt.Println("\thello - end")
}

func main() {
	fmt.Println("main - start")
	defer fmt.Println("This is deferred statement")

	hello()
	os.Exit(7687) // program execution terminates

	// Below code is not reachable
	fmt.Println("main - end")

}
