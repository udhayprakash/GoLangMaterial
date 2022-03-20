package main

import "fmt"

/*Purpose: Unidirectional Channels
send-only 		: chan<- int
receive-only	: <-chan int

*/
func counter(outChannel chan<- int) {
	// outChannel - send-only channel
	for x := 0; x < 10; x++ {
		outChannel <- x
	}
	close(outChannel)
}

func squarer(outChannel chan<- int, inChannel <-chan int) {
	// outChannel - send-only channel
	// inChannel  - receive-only channel
	for val := range inChannel {
		outChannel <- val * val
	}
	close(outChannel)
}

func printer(inChannel <-chan int) {
	// inChannel  - receive-only channel
	for v := range inChannel {
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
