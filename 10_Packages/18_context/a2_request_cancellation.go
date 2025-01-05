//  request cancellation, with context
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go longRunningOperation(ctx)
	time.Sleep(2 * time.Second)
	cancel() // Cancel the operation
	time.Sleep(1 * time.Second)
}
