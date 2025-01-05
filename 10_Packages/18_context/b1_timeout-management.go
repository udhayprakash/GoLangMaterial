// Traditionally request TIMEOUT, without context
package main

import (
	"fmt"
	"time"
)

func queryDatabase() {
	fmt.Println("queryDatabase - started")
	time.Sleep(3 * time.Second) //  if sleep seconds less than 2, no problem
	fmt.Println("Query completed")
}

func main() {
	done := make(chan bool)
	go func() {
		queryDatabase()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Query succeeded")
	case <-time.After(2 * time.Second):
		fmt.Println("Query timed out")
	}
}
