package main

import "fmt"

// func add() {
// 	fmt.Println("called add function")
// }

var add func() = nil

func main() {
	defer add()
	fmt.Printf("value - %v\n", add) // 0xeadd60
	fmt.Printf("Type  - %T\n", add) // func()
}

// panic: runtime error: invalid memory address or nil pointer dereference
