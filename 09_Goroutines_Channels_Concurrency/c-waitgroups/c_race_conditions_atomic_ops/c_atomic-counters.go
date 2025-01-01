package main

/*

The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools.
There are a few other options for managing state though.
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 3; c++ {
				atomic.AddUint64(&ops, 1)
				fmt.Println("within- ops:", ops)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("outside - ops:", ops)
}
