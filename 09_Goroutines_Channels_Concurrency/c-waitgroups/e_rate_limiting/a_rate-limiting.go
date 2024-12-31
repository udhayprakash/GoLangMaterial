package main

/*
Purpose: Rate limiting
	is an important mechanism for controlling
	resource utilization and maintaining quality of service.
	Go elegantly supports rate limiting with goroutines,
	channels, and tickers.
*/
import (
	"fmt"
	"time"
)

func main() {
	// tradition
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i // passing to channel
	}
	close(requests)

	// TRy to receve only in that time period
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests { // receiveing from channel
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// ====================================================
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i // sending
	}
	close(burstyRequests)

	for req := range burstyRequests { // receiving
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
