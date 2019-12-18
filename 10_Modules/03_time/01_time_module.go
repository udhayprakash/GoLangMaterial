package main

import (
	"fmt"
	"time"
)

func main() {
	// To count the number of units in a Duration, divide:
	second := time.Second
	fmt.Println(int64(second/time.Millisecond)) // prints 1000

	// To convert an integer number of units to a Duration, multiply:
	seconds := 10
	fmt.Println(time.Duration(seconds)*time.Second) // prints 10s

	// Sleep pauses the current goroutine for at least the duration d.
	// A negative or zero duration causes Sleep to return immediately.
	fmt.Println("Sleeping ...")
	time.Sleep(100 * time.Millisecond)

	// time duration calculation
	startTime := time.Now()
	time.Sleep(100)
	endTime := time.Now()
	fmt.Printf("The call took %v to run.\n", endTime.Sub(startTime))

	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.\n", m.Seconds())

	t1 := time.Date(2019, time.December, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, time.February, 16, 0, 0, 0, 0, time.UTC)
	fmt.Println(t2.Sub(t1).String()) // 1512h0m0s


	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	fmt.Printf(`year = %v month = %v day = %v`,
		year, month, day)

}