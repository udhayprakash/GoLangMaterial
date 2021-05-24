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
	fmt.Println("starting ...")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, ",") // 9,8,7,6,5,4,3,2,1,0,
	}
	fmt.Println("DONE  !!!")
}
