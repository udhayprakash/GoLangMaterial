package main

import "fmt"

func myFunc() {
	fmt.Println("Inside my goroutine")
}

func main() {
	fmt.Println("Hello World")
	go myFunc()
	fmt.Println("Finished Execution")
}
