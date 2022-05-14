package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for eachJob := range jobs {
		fmt.Printf("Worker:%d - Started  job:%v\n", id, eachJob)
		time.Sleep(time.Second)
		fmt.Printf("\tWorker:%d - Finished job:%v\n", id, eachJob)
		results <- eachJob * 2
	}
}

func main() {
	const workerPoolCount = 100
	jobs := make(chan int, workerPoolCount)
	results := make(chan int, workerPoolCount)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	jobs <- 0
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		// <-results// receiving and discarding values from channel
		fmt.Println("<-results", <-results)
	}
}
