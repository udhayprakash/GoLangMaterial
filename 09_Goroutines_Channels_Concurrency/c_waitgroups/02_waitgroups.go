package main

import (
	"fmt"
	"sync"
	"time"
)
// To wait for multiple goroutines to finish, we can use a wait group.

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worder:%d - Started\n", id)

	time.Sleep(time.Second)
	fmt.Printf("\tWorder:%d - Finished\n", id)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
