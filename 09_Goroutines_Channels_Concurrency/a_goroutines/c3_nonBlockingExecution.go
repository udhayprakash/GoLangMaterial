package main

import (
	"fmt"
	"time"
)

func sleep250() {
	fmt.Println("\tsleep250 - start")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("\tsleep250 - end")
}

func sleep500() {
	fmt.Println("\tsleep500 - start")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\tsleep500 - end")
}

func main() {
	fmt.Println("main -start")
	go sleep250()
	go sleep500()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("main - end")
}

/* OUTPUT - Non-Blocking Execution

- go run c3_blockingExecution.go
main -start
		sleep500 - start
		sleep250 - start
		sleep250 - end
		sleep500 - end
main - end

- go run c3_blockingExecution.go
main -start
		sleep250 - start
		sleep500 - start
		sleep250 - end
		sleep500 - end
main - end

*/
