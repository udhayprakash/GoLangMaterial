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
		for i := 0; i < 30; i++ {
			fmt.Println("one", i)
		}
	}()

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println("two", i)
		}
	}()

	time.Sleep(time.Second)
	elapsedTime := time.Since(start)
	fmt.Println("Total time for execution:", elapsedTime.String())
}

/*
windows command to check the number of cores and logical processors:

	wmic cpu get NumberOfCores,NumberOfLogicalProcessors

		NumberOfCores  NumberOfLogicalProcessors
		4              8

*/
