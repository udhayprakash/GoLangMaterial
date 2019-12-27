package main

import "fmt"

func main() {
	type address struct {
		hostname string
		port int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	fmt.Println(hits)
}
