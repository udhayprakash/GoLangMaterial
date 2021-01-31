package main

import "fmt"

/*
Purpose: Channels
	Types
		- New
		- Unbuffered
		- Buffered
*/

func main(){
	// New Channel
	var c1 chan int // default value is nil
	go myGoRoutine(c1)
	//c1 = 34 //  cannot use 34 (type int) as type chan int in assignment
	//c1 <- 34 // all goroutines are asleep - deadlock!
	//<- c1
	//close(c1) // panic: close of nil channel

	fmt.Printf("c1 - dataType:%T	 Value:%v\n", c1, c1)


	// Unbuffered Channel
	c2 := make(chan int)
	go myGoRoutine(c2)
	// Assigning values to channel
	//c2 <- 11 // fatal error: all goroutines are asleep - deadlock!
	// reading from unbuffered channel
	//<- c2
	//close(c2)
	fmt.Printf("c2 - dataType:%T	 Value:%v\n", c2, c2)


	// Buffered Channel
	c3 := make(chan int, 3) // like array of 3 values
	go myGoRoutine(c3)
	c3 <- 111
	c3 <- 222
	c3 <- 333
	fmt.Printf("c3 - dataType:%T	 Value:%v\n",  c3,  c3)


	// work-around to pause the main function till execution
	var input string
	fmt.Scanln(&input)
}

func myGoRoutine(c chan int) {
	fmt.Printf("\tWORKER:c - dataType:%T	 Value:%v\n",  c,  c)
}