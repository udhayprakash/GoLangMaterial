package main

import (
	"fmt"
	"github.com/dvyukov/go-fuzz/go-fuzz"
)
/*

Fuzz testing, also known as fuzzing, is a technique for testing software by providing it with random or semi-random inputs in order to find vulnerabilities and bugs. In Go, you can use the package "go-fuzz" to perform fuzz testing on your code.

*/

func Fuzz(data []byte) int {
	var x, y int
	if _, err := fmt.Sscanf(string(data), "%d %d", &x, &y); err != nil {
		return 0
	}
	result := add(x, y)
	if result < 0 {
		panic("Negative result")
	}
	return 1
}

func main() {
	fuzz.Run(Fuzz)
}

//  go get github.com/dvyukov/go-fuzz/go-fuzz
// go-fuzz-build && go-fuzz

// TODO - https://go.dev/doc/tutorial/fuzz