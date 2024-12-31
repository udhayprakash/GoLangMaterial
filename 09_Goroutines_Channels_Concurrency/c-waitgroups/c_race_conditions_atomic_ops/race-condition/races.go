package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // using outer context
			wg.Done()
		}()
	}
	wg.Wait()
}

// go run -race races.go
