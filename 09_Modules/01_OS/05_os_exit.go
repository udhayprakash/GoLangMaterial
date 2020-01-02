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

func main() {

	defer fmt.Println("!")

	os.Exit(3)
}
