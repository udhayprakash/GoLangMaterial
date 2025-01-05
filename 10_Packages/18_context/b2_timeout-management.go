// request TIMEOUT, with context
package main

import (
	"context"
	"fmt"
	"time"
)

func queryDatabase(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Query completed")
	case <-ctx.Done():
		fmt.Println("Query cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	queryDatabase(ctx)
}
