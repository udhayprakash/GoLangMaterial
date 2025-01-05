// request cancellation, with context
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func longRunningOperation(ctx context.Context, jobName string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done

	count := 0 // Counter to track the number of times the job runs
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Operation cancelled: %v\n", jobName, ctx.Err())
			fmt.Printf("[%s] Total iterations: %d\n", jobName, count)
			return
		default:
			count++ // Increment the counter
			fmt.Printf("[%s] Working... (Iteration: %d)\n", jobName, count)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Start multiple jobs
	jobs := []string{"job1", "job2", "job3"}
	for _, job := range jobs {
		wg.Add(1) // Increment the WaitGroup counter for each job
		go longRunningOperation(ctx, job, &wg)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Cancelling all jobs...")
	cancel()

	// wait for all goroutines to finish
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")

}
