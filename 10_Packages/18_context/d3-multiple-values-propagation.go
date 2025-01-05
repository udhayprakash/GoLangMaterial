package main

import (
	"context"
	"fmt"
	"time"
)

// Define a custom type for context keys to avoid collisions
// Needed fpor handling multiple data types
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
	isAdminKey   contextKey = "isAdmin"
)

// processRequest processes a request using values from the context
func processRequest(ctx context.Context) {
	// Retrieve values from the context with type assertions
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		fmt.Println("requestID not found or invalid type")
		return
	}

	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		fmt.Println("userID not found or invalid type")
		return
	}

	isAdmin, ok := ctx.Value(isAdminKey).(bool)
	if !ok {
		fmt.Println("isAdmin not found or invalid type")
		return
	}

	// Process the request using the retrieved values
	fmt.Printf("Processing request:\n")
	fmt.Printf("  	Request ID: %s\n", requestID)
	fmt.Printf("  	User      : %d\n", userID)
	fmt.Printf("  	Is Admin  : %v\n", isAdmin)
}

func main() {
	// Create a context with multiple values of different types
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "12345") // string
	ctx = context.WithValue(ctx, userIDKey, 67890)      // int
	ctx = context.WithValue(ctx, isAdminKey, true)      // bool

	// Process the request with the context
	go processRequest(ctx)

	// Wait to see the output
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")
}
