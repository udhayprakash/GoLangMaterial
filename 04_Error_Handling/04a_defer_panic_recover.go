package main

import "fmt"

func helloWorld()  {
	fmt.Println("Hello World")
	fmt.Println("recover() :", recover())
}
func main() {
	fmt.Println("main func start")
	defer helloWorld()
	panic("PANIC")
	fmt.Println("main func end")
}
// Ensure that "panic" follows after "defer"
// and recover() is used in deferred function

// Statements after panic() are not reachable