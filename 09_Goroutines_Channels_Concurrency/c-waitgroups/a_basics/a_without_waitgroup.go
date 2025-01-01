package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	done <- true
}

func main() {
	done := make(chan bool)
	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}
	for i := 1; i <= 3; i++ {
		<-done // earlier solution
	}
	fmt.Println("All workers done")
}
