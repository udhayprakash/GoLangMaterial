package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// when we want to wait till
	until, _ := time.Parse(time.RFC3339, "2023-06-22T15:04:05+02:00")

	// time.Sleep(time.Until(until))

	waitUntil(ctx, until)

}

func waitUntil(ctx context.Context, until time.Time) {
	fmt.Println("waitUntil - start")
	timer := time.NewTimer(time.Until(until))
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("waitUntil - <-timer.C")
		return
	case <-ctx.Done():
		fmt.Println("waitUntil - <-ctx.Done()")
		return
	}
}
