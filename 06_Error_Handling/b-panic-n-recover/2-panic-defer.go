package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println("This is deferred1")
	defer fmt.Println("This is deferred2")
	panic("something bad happened")

	fmt.Println("end")

}

/*
NOTE:
	1. panic() is the last executing statement
	2. After panicking, it will execute all registered defers, till that line
*/

// start
// This is deferred2
// This is deferred1
// panic: something bad happened
