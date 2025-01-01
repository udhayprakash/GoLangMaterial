package main

import (
	"fmt"
	"time"
)

func main() {
	// time channel
	timer1 := time.NewTimer(5 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
