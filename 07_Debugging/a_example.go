package main

import "fmt"

func main() {
	msg := "Hello, Delve!"
	fmt.Println(msg)
}

/*
NOTE:
	To know current installed delve version,
		dlv version

	Debugging
		go run a_example.go
		dlv debug a_example.go

	To set breakpoint,
		break a_example.go:7

	To continue execution till breakpoint,
		continue

	All commands are similar to that of pdb of python.
*/
