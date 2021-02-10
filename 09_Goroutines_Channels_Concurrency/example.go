package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// allowed to be executed on multiple cores parallelly,
	// ensuring faster executions
	runtime.GOMAXPROCS(4)

	start := time.Now()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("one", i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("two", i)
		}
	}()

	elapsedTime := time.Since(start)
	fmt.Println("Total time for execution:", elapsedTime.String())
	time.Sleep(time.Second)
}
