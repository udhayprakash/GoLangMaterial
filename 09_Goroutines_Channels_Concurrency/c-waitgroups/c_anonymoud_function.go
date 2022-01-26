package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine")
		waitgroup.Done()
	}()
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}
