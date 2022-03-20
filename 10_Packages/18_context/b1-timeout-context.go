package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := make(chan string)
	go func() {
		var answer string
		fmt.Print("Enter something: ")
		fmt.Scanln(&answer)
		result <- answer
	}()

	select {
	case r := <-result:
		fmt.Println("answer is      :", r)
	case <-ctx.Done():
		fmt.Println("timeout")
	}

}

// assignment: implment number guessing game to work for 1 minute, only. return scores
