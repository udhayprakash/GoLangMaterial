package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	go func() {
		for {
			x++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("x =", x) // x = 0, why?
}
