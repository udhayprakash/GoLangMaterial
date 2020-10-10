package main
/*
A defer statement defers the execution of a function until
the surrounding function returns.

The deferred call's arguments are evaluated immediately,
but the function call is not executed until the surrounding
function returns.
Deferred functions are run even if a runtime panic occurs.

Go’s garbage collector recycles unused memory, but do not assume
it will release unused operating system resources like open files
and network connections. They should be closed explicitly
*/
import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
