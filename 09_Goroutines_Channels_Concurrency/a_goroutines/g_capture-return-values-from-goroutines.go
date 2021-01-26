package main

import "fmt"

func say(msg chan<- string) {
	msg <- "Hello, 世界!"
}

// return multiple values, in this case it is 2
func trueOrFalse(answer chan<- bool) {
	answer <- true
	answer <- false
}

func main() {
	messageChannel := make(chan string)

	// execute Gorountine and return the value into messageChannel
	go say(messageChannel)

	fmt.Println(<-messageChannel)

	// the channel has type. To receive 2 boolean values, we use increase it by 2

	booleanChannel := make(chan bool, 2)

	go trueOrFalse(booleanChannel)

	fmt.Println(<-booleanChannel) // true
	fmt.Println(<-booleanChannel) // false -- and also pop the stack

}
