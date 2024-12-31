package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched() means let the CPU execute
		// other goroutines, and come back at some point.
		runtime.Gosched()
		fmt.Println(s, i)

	}
}

func main() {
	// Set the number of CPU cores that Go will use for concurrency
	// runtime.GOMAXPROCS(2) // This tells Go to use 2 CPU cores

	// Check the number of CPU cores Go will use
	fmt.Println("Go is using", runtime.NumCPU(), "CPU cores")
	
	go say("\tworld") // create a new goroutine
	say("hello")      // current goroutine
}

/*
The scheduler only uses one thread to run all goroutines, which means it only implements
concurrency. If you want to use more CPU cores in order to take advantage of parallel processing,
you have to call runtime.GOMAXPROCS(n) to set the number of cores you want to use. If n<1,
it changes nothing.
By default,
	before 1.5 version, GOMAXPROCS = 1 &
	from 1.5 release, GOMAXPROCS = number of cores

Scheduling:

The Go runtime schedules goroutines on available OS threads using a mechanism called the M:N scheduler (many goroutines mapped to many threads). Some runtime goroutines manage this scheduling.

*/
