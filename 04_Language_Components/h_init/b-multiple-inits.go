package main

import "fmt"

func init() {
	fmt.Println("First init")
}

func init() {
	fmt.Println("Second init")
}

func initThird() { // Not an init function, but ordinary function
	fmt.Println("Third init")
}

func init() {
	fmt.Println("Fourth init")
}

func Init() {
	fmt.Println("Fifth init")
}
func main() {}

// NOTE: init() functions will execute in order of definition
