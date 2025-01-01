package main

import (
	"fmt"
	"sync"
	"time"
)

func longConcurrentProcess(sleep int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Sleeping for", sleep, "seconds ...")
	time.Sleep(time.Duration(sleep) * time.Second)
}

func main() {
	var wg sync.WaitGroup

	total := 5
	// wg.Add(total)
	// wg.Add(total+2)  // fatal error: all goroutines are asleep - deadlock!

	for i := 1; i <= total; i++ {
		go longConcurrentProcess(i, &wg)
	}

	// Wait for all goroutines to be finished
	wg.Wait()
	fmt.Println("Finished")

}
