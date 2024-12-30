package main

import (
	"fmt"
	"time"
)

/*

Blocking
	statement1  -- 5 sec
	statement2  -- 3 sec
	statement3  -- 3 sec
	--------------------
			TOTAL: 5 + 3 + 3 = 11 sec


Non-Blocking Execution    --> go routine
	statement1  -- 5 sec
	statement2  -- 3 sec
	statement3  -- 3 sec
	--------------------
			TOTAL:             5 sec

*/

func printItem(i int) {
	fmt.Printf("\tPrint Item: %2v\n", i+1)
}

func main() {
	fmt.Println("Start Script")

	for j := 0; j < 10; j++ {
		go printItem(j)
	}
	fmt.Println("End Script")

	// Wait, giving time for the go routines to finish.
	// time.Sleep(10000000)
	time.Sleep(time.Second * 2)
}

// NOTE: Go-routines may not execute in the order of initialization

/*
$ go run b_blockingExecution.go 
Start Script
        Print Item:  1
        Print Item:  2
        Print Item:  3
        Print Item:  4
        Print Item:  5
        Print Item:  6
        Print Item:  7
        Print Item:  8
        Print Item:  9
        Print Item: 10
End Script

$ go run b_blockingExecution.go 
Start Script
End Script
        Print Item: 10
        Print Item:  1
        Print Item:  2
        Print Item:  3
        Print Item:  4
        Print Item:  5
        Print Item:  6
        Print Item:  7
        Print Item:  8
        Print Item:  9


*/
