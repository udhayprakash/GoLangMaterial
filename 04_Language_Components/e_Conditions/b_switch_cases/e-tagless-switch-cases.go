package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("t.Hour() = ", t.Hour())

	// Tagless switch
	switch {
	case t.Hour() <= 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
