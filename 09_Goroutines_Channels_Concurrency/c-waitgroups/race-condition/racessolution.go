package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			wg.Done()
		}
	}()
	wg.Wait()
}

// go run -race races.go
