package main

import "fmt"

// Golang doesnt support default args
// declaration directly in definition

// Method 1
func demo(greeting string) {
	fmt.Println(greeting)
}

func demoWithDefault() {
	demo("hello")
}

//  Method 2
func demo2(g ...string) {
	// fmt.Println(g, len(g))
	switch len(g) {
	case 0:
		fmt.Println("Hello")
	case 1:
		fmt.Println(g[0])
	case 2:
		fmt.Println(g[0:2])
	default:
		panic("too many arguments")
	}

}

func main() {
	demo("Good Morning")
	// demo() // not enough arguments in call to demo

	demoWithDefault()

	demo2()
	demo2("Hello")
	demo2("Hello", "world")
	// demo2("Hello", "world", "87") // panic: too many arguments

}
