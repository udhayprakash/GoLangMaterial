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

// Method 2
func demo(g ...string) {
	switch len(g) {
	case 0:
		fmt.Println("Hello")
	case 1:
		fmt.Println(g[0])
	default:
		panic("too many arguments")
	}
}

// Method 3
func demo(g *string) {
	if g == nil {
		g = "Hello"
	}

	fmt.Println(g)
}
