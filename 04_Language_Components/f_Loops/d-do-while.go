package main

import (
	"fmt"
)

/*
Go Doesnt support do-while loop
*/
func main() {
	for ok := true; ok; ok = false {
		fmt.Println("IN Loop1")
	}

	for ok := true; ok; {
		fmt.Println("IN Loop2")
		ok = false
	}

	ok := true
	for ok {
		fmt.Println("IN Loop2")
		ok = false
	}
}
