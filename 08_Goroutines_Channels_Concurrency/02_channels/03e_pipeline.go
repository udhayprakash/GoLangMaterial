package main

import (
	"fmt"
)

/*Purpose: Unidirectional Channels
send-only 		: chan<- int
receive-only	: <-chan int

*/
func counter(out chan<- int) {
	// out - send-only channel
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	// out - send-only channel
	// in  - receive-only channel
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	// in  - receive-only channel
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)

	printer(squares)
}
