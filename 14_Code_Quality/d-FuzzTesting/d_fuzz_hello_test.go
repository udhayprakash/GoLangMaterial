package main

import (
	"fmt"
	"github.com/dvyukov/go-fuzz/go-fuzz"
)

func Fuzz(data []byte) int {
	name := string(data)
	result := Hello(name)
	if result == "" {
		return 0
	}
	return 1
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fuzz.Run(Fuzz)
}

// go get -u github.com/dvyukov/go-fuzz/...