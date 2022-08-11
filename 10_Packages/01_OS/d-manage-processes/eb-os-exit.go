package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("main() - started")

	// defer function execution
	defer func() {
		fmt.Println("main completed!")
	}()

	// run function after 2 seconds
	time.AfterFunc(2*time.Second, func() {
		os.Exit(1) // exit with status code 1
	})

	// sleep goroutine for 3 seconds
	time.Sleep(3 * time.Second)
}
