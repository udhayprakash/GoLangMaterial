// we can achieve syncronization between go routines by using a
//  channel but only sending back a bool to say it’s complete.
package main

import (
	"fmt"
	"time"
)

func main() {

	isDone := make(chan bool, 1)

	go work(isDone)

	<-isDone
	fmt.Println("Finished")
}

func work(isDone chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	isDone <- true
}
