package main

/*
Purpose: Go-Routines
	A goroutine is a lightweight thread of execution.

	non-blocking execution
*/
import (
	"fmt"
	"time"
)

func f1(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f1("direct")

	go f1("goroutine")

	go func(msg string) { // anonymous function, or, function literal
		fmt.Println("msg=", msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
