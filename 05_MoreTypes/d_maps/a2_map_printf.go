package main

import "fmt"

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func main() {
	// For maps, Printf and friends sort the output 
	// lexicographically by key.
	fmt.Printf("%#v\n", timeZone)

}
