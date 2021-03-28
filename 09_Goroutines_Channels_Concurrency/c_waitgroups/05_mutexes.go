package main

import (
	"fmt"
	"sync"
)

// mutex - needed for solving race condition
func main() {
	m := make(map[int]string)
	m[2] = "First Value"

	var lock sync.Mutex
	go func() {
		lock.Lock()
		m[2] = "Second Value"
		lock.Unlock()
	}()

	lock.Lock()
	v := m[2]

	lock.Unlock()
	fmt.Println(v)
}
