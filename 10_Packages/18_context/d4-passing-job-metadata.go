package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate a background job
func processJob(ctx context.Context) {
	jobID := ctx.Value("jobID").(string)
	retryCount := ctx.Value("retryCount").(int)
	
	fmt.Printf("[JobID: %s] Processing job (Retry: %d)...\n", jobID, retryCount)
	
	time.Sleep(1 * time.Second)
	fmt.Printf("[JobID: %s] Job processed.\n", jobID)
}

func main() {
	// Create a context with job metadata
	ctx := context.WithValue(context.Background(), "jobID", "job-123")
	ctx = context.WithValue(ctx, "retryCount", 3)

	// Process the job
	go processJob(ctx)

	// Simulate other work
	time.Sleep(2 * time.Second)
	fmt.Println("Main function exiting.")
}
