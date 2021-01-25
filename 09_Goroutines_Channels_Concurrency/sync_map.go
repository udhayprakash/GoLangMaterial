package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("example", []int{1, 2, 3})
	ints, found := m.Load("example")
	fmt.Println(ints, found)
}
