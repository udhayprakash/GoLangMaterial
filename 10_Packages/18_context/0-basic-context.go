package main

import (
	"context"
	"fmt"
	"reflect"
)

func InspectContext(ct context.Context) {
	/*
	context.Context Methods:
			Deadline()	: Gets the context’s deadline.
			Done()		: Signals cancellation via a channel.
			Err()		: Returns the cancellation reason.
			Value()		: Retrieves a value from the context.
	*/
	
	// Inspect the type of the context
	fmt.Println("Type of ctx:", reflect.TypeOf(ct))

	// Inspect the methods of the context
	for i := 0; i < reflect.TypeOf(ct).NumMethod(); i++ {
		method := reflect.TypeOf(ct).Method(i)
		fmt.Printf("\tMethod %d: %s\n", i+1, method.Name)
	}

}


func main() {
	ctx := context.Background()
	fmt.Println("\nUsing background context:", ctx)
	InspectContext(ctx)

	ctx = context.TODO()
	fmt.Println("\nUsing TODO context:", ctx)
	InspectContext(ctx)

	// Create a context with a value
	ctx = context.WithValue(context.Background(), "userID", 12345)
	fmt.Println("\nUsing WithValue context:", ctx)
	InspectContext(ctx)

	// Access the value in the context
	userID := ctx.Value("userID")
	fmt.Println("User ID:", userID)
}
