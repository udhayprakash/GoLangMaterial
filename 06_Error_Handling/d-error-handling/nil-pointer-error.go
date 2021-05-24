package main

import "fmt"

var add func() = nil

// func add() {
// 	fmt.Println("called add function")
// }

func main() {
	fmt.Printf("value - %v\n", add)
	fmt.Printf("Type  - %T\n", add)
	defer add()
}

// invalid memory address or nil pointer dereference
