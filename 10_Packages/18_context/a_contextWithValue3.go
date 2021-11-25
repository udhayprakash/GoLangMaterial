package main

import (
	"context"
	"log"
	"time"
)

func cron(ctx context.Context, startTime time.Time, delay time.Duration) <-chan time.Time {
	// Create the channel which we will return
	stream := make(chan time.Time, 1)

	// Calculating the first start time in the future
	// Need to check if the time is zero (e.g. if time.Time{} was used)
	if !startTime.IsZero() {
		diff := time.Until(startTime)
		if diff < 0 {
			total := diff - delay
			times := total / delay * -1

			startTime = startTime.Add(times * delay)
		}
	}

	// Run this in a goroutine, or our function will block until the first event
	go func() {

		// Run the first event after it gets to the start time
		t := <-time.After(time.Until(startTime))
		stream <- t

		// Open a new ticker
		ticker := time.NewTicker(delay)
		// Make sure to stop the ticker when we're done
		defer ticker.Stop()

		// Listen on both the ticker and the context done channel to know when to stop
		for {
			select {
			case t2 := <-ticker.C:
				stream <- t2
			case <-ctx.Done():
				close(stream)
				return
			}
		}
	}()

	return stream
}

func main() {
	// Run on Tuesdays by 2 pm =========================
	ctx := context.Background()

	startTime, err := time.Parse(
		"2006-01-02 15:04:05",
		"2019-09-17 14:00:00",
	) // is a tuesday
	if err != nil {
		panic(err)
	}

	delay := time.Hour * 24 * 7 // 1 week

	for t := range cron(ctx, startTime, delay) {
		// Perform action here
		log.Println(t.Format("2006-01-02 15:04:05"))
	}

	// Run every hour, on the hour ======================
	ctx2 := context.Background()

	startTime2, err := time.Parse(
		"2006-01-02 15:04:05",
		"2019-09-17 14:00:00",
	) // any time in the past works but it should be on the hour
	if err != nil {
		panic(err)
	}

	delay2 := time.Hour // 1 hour

	for t := range cron(ctx2, startTime2, delay2) {
		// Perform action here
		log.Println(t.Format("2006-01-02 15:04:05"))
	}

	// Run every 10 minutes, starting in a week =====================
	ctx3 := context.Background()

	startTime3, err := time.Now().AddDate(0, 0, 7) 
	// see https://golang.org/pkg/time/#Time.AddDate
	if err != nil {
		panic(err)
	}

	delay3 := time.Minute * 10 // 10 minutes

	for t := range cron(ctx3, startTime3, delay3) {
		// Perform action here
		log.Println(t.Format("2006-01-02 15:04:05"))
	}

}
