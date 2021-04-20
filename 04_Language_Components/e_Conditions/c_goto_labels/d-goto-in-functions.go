package main

import "fmt"

func main() {
	// goto GlobalLabel
	learnGoTO()
}

func learnGoTO() {
	fmt.Println("A")
	goto FINISH
	fmt.Println("B")
}

func test() {
FINISH:
	fmt.Println("C")
}

// GLobalLabel:
// 	fmt.Println("IN GLobalLabel")

/*
NOTE:
	1. Label need to be defined witin functions only
	2. goto labels have scope within same function, as defined

*/
