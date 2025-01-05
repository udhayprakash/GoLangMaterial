// solution with context
package main

import (
	"fmt"
	"time"
)

func processRequest(ctx context.Context) {
    requestID := ctx.Value("requestID").(string)
    fmt.Println("Processing request with ID:", requestID)
}

func main() {
    ctx := context.WithValue(context.Background(), "requestID", "12345")
    go processRequest(ctx)
    time.Sleep(1 * time.Second) // Wait to see the output
}