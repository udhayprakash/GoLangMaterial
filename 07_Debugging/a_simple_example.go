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
		go run a_simple_example.go
		dlv debug a_simple_example.go

	To set breakpoint,
		break a_simple_example.go:7

	To continue execution till breakpoint,
		continue

	NOTE: All commands are similar to that of pdb of python.

When compiling, always disable optimizations and include debug symbols:

	go build -gcflags="all=-N -l"

*/