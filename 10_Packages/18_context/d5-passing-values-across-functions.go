/*
Passing context values across functions

	some read only
	some read and write too


	main() -> handleRequest
				logRequest  	 (read only from context)
				authenticateUser (reads and writes value, from/to context); isAdminKey
				processRequest	 (read only from context)
*/	
package main

import (
	"context"
	"fmt"
	"time"
)

// Define custom types for context keys to avoid collisions
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
	authTokenKey contextKey = "authToken"
	isAdminKey   contextKey = "isAdmin"
)

func logRequest(ctx context.Context) {
	requestID := ctx.Value(requestIDKey).(string)
	userID := ctx.Value(userIDKey).(int)
	fmt.Printf("[Log] Request ID: %s, User ID: %d\n", requestID, userID)
}

func authenticateUser(ctx context.Context) context.Context {
	authToken := ctx.Value(authTokenKey).(string)
	fmt.Printf("[Auth] Validating token: %s\n", authToken)

	if authToken == "valid-token" {
		fmt.Println("[Auth] User authenticated successfully.")
		// Add isAdmin flag to the context
		ctx = context.WithValue(ctx, isAdminKey, true)
	} else {
		fmt.Println("[Auth] Invalid token. Authentication failed.")
		ctx = context.WithValue(ctx, isAdminKey, false)
	}
	return ctx
}

func processRequest(ctx context.Context) {
	requestID := ctx.Value(requestIDKey).(string)
	userID := ctx.Value(userIDKey).(int)
	isAdmin := ctx.Value(isAdminKey).(bool)
	fmt.Printf("[Process] Processing request ID: %s for user ID: %d\n", requestID, userID)

	// Check if the user is an admin
	if isAdmin {
		fmt.Println("[Process] User is an admin. Granting full access.")
	} else {
		fmt.Println("[Process] User is not an admin. Granting limited access.")
	}

	// Simulate processing
	time.Sleep(1 * time.Second)
	fmt.Println("[Process] Request processed successfully.")
}

func handleRequest(ctx context.Context) {
	logRequest(ctx)

	ctx = authenticateUser(ctx)

	processRequest(ctx)
}

func main() {
	// Simulate a request with context values
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-12345")   // Add request ID
	ctx = context.WithValue(ctx, userIDKey, 67890)            // Add user ID
	ctx = context.WithValue(ctx, authTokenKey, "valid-token") // Add auth token

	// Handle the request
	handleRequest(ctx)

	// Simulate another request with an invalid token
	fmt.Println("\nSimulating another request with an invalid token...")
	ctx = context.WithValue(ctx, authTokenKey, "invalid-token") // Update auth token
	handleRequest(ctx)
}
