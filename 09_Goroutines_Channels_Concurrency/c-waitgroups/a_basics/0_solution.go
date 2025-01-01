package main

import (
	"fmt"
	"sync"
)

/*

Method1 -- time sleep
Method2 -- fmt.Scanf
Method3 --- <-channel
Method4 -- waitgroups

*/

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}

func main() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup

	
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	

	waitgroup.Wait()
	fmt.Println("Finished Execution")
}
