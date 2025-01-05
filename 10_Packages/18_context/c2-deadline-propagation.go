// solution with context

package main

import (
	"context"
	"fmt"
	"time"
)

func processTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	go processTask(ctx)
	time.Sleep(3 * time.Second) // Wait to see the output
}
