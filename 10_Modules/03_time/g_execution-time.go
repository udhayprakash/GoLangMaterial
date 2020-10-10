package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(10)

	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println("duration              :", duration)

	// Nanoseconds as int64
	fmt.Println("duration.Nanoseconds():", duration.Nanoseconds())

	// Alternatively
	end := time.Now()
	duration = end.Sub(start)
	fmt.Println("\nduration              :", duration)

	// Measure a function call
	foo()
}

func foo() {
	defer duration(track("foo"))
	// Code to measure
	i := 1
	for i <= 10 {
		i++
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
