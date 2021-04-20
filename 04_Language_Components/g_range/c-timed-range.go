package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for range time.Tick(time.Second * 2) { // runs for every 2 seconds
		fmt.Printf("Hello-%d\n", count)
		if count > 5 {
			break
		}
		count += 1 // count++
	}
}
