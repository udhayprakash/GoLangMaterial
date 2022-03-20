package main

import (
	_ "fmt"
)

func main() {
	// The _ operator actually means
	// we just want to import that package and execute its init function,
	// and we are not sure if want to use the functions belonging to that package.

	// init() in `fmt` package is executed.
}
