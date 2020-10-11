package main

import (
	"fmt"
	"time"
)

var weekday string

func init() {
	fmt.Println("Init is called ...")
	weekday = time.Now().Weekday().String()
}

func main() {
	fmt.Println("main is called ...")
	fmt.Printf("Today is %s", weekday)
}

/*
The main purpose of the init() function is to initialize the global variables
that cannot be initialized in the global context.

init() can be used in all packages, whereas main() can only be used in main package.
It is recommended to use only one init() per package
Go programs will call main() and init() automatically
Both main() and init() will not take any arguments and wornt return any
*/