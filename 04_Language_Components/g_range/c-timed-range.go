package main

import (
	"fmt"
	"time"
)

func main() {
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
 