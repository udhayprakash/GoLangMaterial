package main

import "fmt"

func main() {
	fmt.Println("start")

	defer ErrorHandler()
	panic("something bad happened")

	fmt.Println("end")

}

func ErrorHandler() {
	fmt.Println("In Error Handler func")
	// fmt.Println(recover())  // something bad happened

	if err := recover(); err != nil {
		fmt.Println("Error is:", err)
	}
}

/*

NOTE - defer executed before panic
Just calling recover() func in defer func
will stop panicing
recover() is last to execute

recover() is useful only in defer functions


Ensure that "panic" follows after "defer"
and recover() is used in deferred function

Statements after panic() are not reachable
*/
 