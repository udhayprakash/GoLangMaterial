package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched() means let the CPU execute other goroutines, and come back at some point.
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("\tworld") // create a new goroutine
	say("hello")      // current goroutine
}
/*
The scheduler only uses one thread to run all goroutines, which means it only implements
concurrency. If you want to use more CPU cores in order to take advantage of parallel processing,
you have to call runtime.GOMAXPROCS(n) to set the number of cores you want to use. If n<1,
it changes nothing.
Ref: http://concur.rspace.googlecode.com/hg/talk/concur.html#landing-slide

*/