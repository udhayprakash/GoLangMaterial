package main

import (
	"context"
	"fmt"
	"time"
)

func nested2(ctx context.Context) {
	contextWithTimeout, cancelFunction := context.WithTimeout(ctx, 2*time.Second)
	select {
	case <-contextWithTimeout.Done():
		fmt.Println("Nested2 finished execution")
		cancelFunction()
	case <-ctx.Done():
		fmt.Println("Parent got terminated")
		cancelFunction()
	}
}

func nested1(ctx context.Context) {
	contextWithTimeout, cancelFunction := context.WithTimeout(ctx, 3*time.Second)
	nested2(contextWithTimeout)
	select {
	case <-contextWithTimeout.Done():
		fmt.Println("COntext done after 3s timeout")
		cancelFunction()
	case <-ctx.Done():
		fmt.Println("Main COntext is closed")
		cancelFunction()
	}
}
func main() {
	ctx := context.Background()
	contextwithcancel, cancelFunction := context.WithCancel(ctx)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling function main")
		cancelFunction()
	}()
	defer func() {
		fmt.Println("Cancel function called main")
		cancelFunction()
	}()
	nested1(contextwithcancel)
}
