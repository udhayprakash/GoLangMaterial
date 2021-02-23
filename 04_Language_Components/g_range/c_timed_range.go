package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for range time.Tick(time.Second * 2){ // runs for every 2 seconds
		fmt.Println(count, "Hello, world")
		count += 1
		if count > 5{
			break
		}
	}
}


