package main

import (
	"log"
	"time"
)

func main() {
	LongRunningFunction()
}

func LongRunningFunction() {
	// Use our time taken function to log the current time with the defer
	// so it will actually run at the end at the end of this function.
	defer TimeTaken(time.Now(), "LongRunningFunction")

	// ... to illustrate pause.
	time.Sleep(2 * time.Second)

}

func TimeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s took %s\n", name, elapsed)
}
