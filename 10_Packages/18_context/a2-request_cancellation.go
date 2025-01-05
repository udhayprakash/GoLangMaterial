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
			fmt.Printf("[%s] Operation cancelled: %v\n", jobName, ctx.Err())
			return
		default:
			fmt.Printf("[%s] Working...\n", jobName)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go longRunningOperation(ctx, "job1")
	go longRunningOperation(ctx, "job2")

	time.Sleep(2 * time.Second)
	cancel() // Cancel the operation
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")

}
