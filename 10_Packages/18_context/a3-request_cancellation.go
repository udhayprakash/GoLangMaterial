// request cancellation, with context
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context, jobName string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%v %s] Operation cancelled: %v\n",
				time.Now().Format("15:04:05"), jobName, ctx.Err())
			return
		default:
			fmt.Printf("[%v %s] Working...\n", time.Now().Format("15:04:05"), jobName)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx := context.TODO()
	go longRunningOperation(ctx, "job1-TODO")
	go longRunningOperation(ctx, "job2-TODO")

	// Simulate some work
	time.Sleep(2 * time.Second)

	fmt.Println("\n\nReplacing TODO context with a cancellable context...")
	ctx, cancel := context.WithCancel(context.Background())

	// restart goroutines with new context
	go longRunningOperation(ctx, "job1-WithCancel")
	go longRunningOperation(ctx, "job2-WithCancel")

	time.Sleep(2 * time.Second)
	cancel() // Cancel the operation
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")

}
