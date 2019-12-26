package main // package declaration

import "fmt"

func main() {
	fmt.Println("Hello, world") // To print a line
}

//Q) How to run a go lang program ?
//Ans)  go run 01_hello.go

//Q) How to build a go lang program ?
//Ans)  go build 01_hello.go
// Then, observe that a new .exe file will be created with the same name

// Package main is special. It defines a standalone executable program, not a library.
// Within package main the function main is also special—it’s where execution of the program begins.
// Whatever main does is what the program does.

// import declarations must follow the package declaration.

// Go does not require semicolons at the ends of statements or declarations,
// except where two or more app ear on the same line.
