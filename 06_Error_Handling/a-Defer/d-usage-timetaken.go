package main

import (
	"fmt"
	"time"
)


func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	fmt.Printf("\tTIME: %s took %s\n", name, elapsed)
}


func main() {
	longRunningFunction()
}

func longRunningFunction() {
	fmt.Println("longRunningFunction - start")

	// Use our time taken function to log the current time with the defer
	// so it will actually run at the end at the end of this function.
	defer timeTaken(time.Now(), "longRunningFunction")

	// ... to illustrate pause.
	time.Sleep(2 * time.Second)
	fmt.Println("longRunningFunction - end")
}
