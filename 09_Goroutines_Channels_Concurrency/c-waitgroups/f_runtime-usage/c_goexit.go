package main

import (
	"fmt"
	"runtime"
	"sync"
)

// runtime.Goexit() - it will stop a goroutine, in which it is running

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go f2("F2", &wg)
	fmt.Println("Main: Waiting for Goroutines to finish")

	wg.Wait()
	fmt.Println("Main completed")
}

func f2(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 0; index < 10; index++ {
		if index == 5 {
			runtime.Goexit()
		}
		fmt.Printf("%v: index %d\n", name, index)
	}
}
