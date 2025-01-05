// Traditional solution
// Example: Passing request-scoped data like a user ID or trace ID.

package main

import (
	"fmt"
	"time"
)

func processRequest(requestID string) {
	fmt.Println("Processing request with ID:", requestID)
}

func main() {
	requestID := "12345"
	go processRequest(requestID)
	time.Sleep(1 * time.Second) // Wait to see the output
}
