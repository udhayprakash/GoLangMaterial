//go:build !test
// +build !test

// empty line after build tag line is mandatory
package main

import (
	"fmt"
)

func Hello() string {
	return "Hello World"
}

func main() {
	fmt.Println(Hello())
}
