package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("This is deferred")
	panic("something bad happened")
	fmt.Println("end")

}
// NOTE - defer executed before panic
