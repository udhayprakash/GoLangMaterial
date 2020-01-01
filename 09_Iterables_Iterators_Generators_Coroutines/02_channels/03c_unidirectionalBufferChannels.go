package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		//for {
		//	//x := <- naturals
		//	x, ok := <-naturals
		//	if !ok {
		//		break
		//	}
		//	squares <- x * x
		//}

		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
