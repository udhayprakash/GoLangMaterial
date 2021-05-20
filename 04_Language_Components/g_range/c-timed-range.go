package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	// runs for every 2 seconds
	for range time.Tick(time.Second * 2) {
		count++
		fmt.Printf("Loop - %d\n", count)

		if count == 5 {
			break
		}
	}
}
