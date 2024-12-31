package main

import (
	"fmt"
	"sync"
)


func increment(wg *sync.WaitGroup, m *sync.Mutex, x1 *int) {
	defer wg.Done()

	m.Lock()
	*x1 = *x1 + 1
	m.Unlock()

}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	var x = 0

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m, &x)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
