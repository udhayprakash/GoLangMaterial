package main

import "fmt"

func main() {
	fmt.Println("In main() loop - start")

	myFunc()
}

func myFunc() {
	fmt.Println("In myFunc() function - start ")
}

func init() {
	fmt.Println("In init() function - start ")
}

/*
The main purpose of the init() function is to initialize the global variables
that cannot be initialized in the global context.

init() can be used in all packages, whereas main() can only be used in main package.
It is recommended to use only one init() per package
Go programs will call main() and init() automatically
Both main() and init() will not take any arguments and wornt return any
*/
