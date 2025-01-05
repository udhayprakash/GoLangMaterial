// solution with context
package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Task %d running...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 3; i++ {
		go task(ctx, i)
	}
	time.Sleep(2 * time.Second)
	cancel() // Stop all tasks
	time.Sleep(1 * time.Second)
}
