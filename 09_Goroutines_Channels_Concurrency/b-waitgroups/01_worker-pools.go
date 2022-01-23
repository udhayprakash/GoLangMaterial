package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for eachJob := range jobs {
		fmt.Printf("Worder:%d - Started  job:%v\n", id, eachJob)
		time.Sleep(time.Second)
		fmt.Printf("\tWorder:%d - Finished job:%v\n", id, eachJob)
		results <- eachJob * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
