package main

import (
	"fmt"
	"time"
)

// run loop  for every 2 seconds
func main() {

	// Infinite Loop
	// for range time.Tick(time.Second * 2) {
	// 	fmt.Println("HElllo")
	// }

	count := 0

	// run loop  for every 2 seconds
	for t := range time.Tick(time.Second * 2) {
		count++
		fmt.Println("Loop -", count, t)

		if count == 5 {
			break
		}
	}
}
