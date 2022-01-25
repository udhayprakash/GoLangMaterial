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
	time.Sleep(1000000)
}

// NOTE: Observe the go-routines may not execute in
//       the order of initialization
