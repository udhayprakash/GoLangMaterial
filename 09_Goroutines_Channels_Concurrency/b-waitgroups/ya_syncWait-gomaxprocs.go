package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
The Go runtime runs goroutines within a logical processor. One logical
processor is bound to one operating system thread. We can specify how many
logical processors will be used in a program. We can set a number of logical
processors by setting the value to runtime.GOMAXPROCS() function. The
default value for GOMAXPROCS is 1. It means that multiple goroutines can be
mapped on an OS thread.
Adding more values to GOMAXPROCS means adding more logical processors. If
we run a program where GOMAXPROCS is 2, and the machine has multiple
core/processors, then the program will run parallely. If the machine has a
single-core, then the program will run concurrently using multiple threads
*/

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go f1("F1", &wg)
	go f1("F2", &wg)
	fmt.Println("Main: Waiting for Goroutines to finish")

	wg.Wait()
	fmt.Println("Main completed")
}

func f1(name string, wg *sync.WaitGroup) {
	for index := 0; index < 10; index++ {
		fmt.Printf("%v: index %d\n", name, index)
	}
	wg.Done()
}
