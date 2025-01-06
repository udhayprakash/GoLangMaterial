package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// // Defer a function to recover from panics
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		// Log the panic and continue execution
	// 		glog.Errorf("Recovered from panic: %v", r)
	// 	}
	// }()
	// panic("This is a simulated panic")


	// Log a fatal message
	glog.Fatal("This is a fatal log. The program will exit.")

	// This line will not be executed
	glog.Info("This log will not be printed")
}

/*
	$ go run e-fatal-logging.go -logtostderr=true

	Program terminates (exit status 1)
	Displays stack trace , to help debug issue
		shows sequence of function calls leading up to the crash

	Runtime information
		about the program's state at the time of the crash, such as:
			Goroutine IDs and their states.
			CPU register values.
			Memory addresses


NOTE: panic() can be recovered using recover()
	But not glog.Fatal
*/
