package main

/*
Purpose:
	- Closing a channel indicates that no more values will be sent on it.
	- This can be useful to communicate completion to the channel’s receivers.

*/
import "fmt"

func main() {
	jobs := make(chan int, 5) // buffered channel
	done := make(chan bool)   // unbuffered channel

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)  // -- WE NEED TO ENSURE CLOSING CHANNEL AT THE END
	//jobs <- 765765765 // panic: send on closed channel
	fmt.Println("sent all jobs")

	<-done

}