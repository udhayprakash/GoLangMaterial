package main

/*
Stacking Defers
	- Deferred function calls are pushed onto a stack.
	- When a function returns, its deferred calls are executed
      in last-in-first-out order.
ref: https://blog.golang.org/defer-panic-and-recover
*/
import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
