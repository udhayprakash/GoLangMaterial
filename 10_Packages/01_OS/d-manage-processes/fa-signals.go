package main

/*
Purpose: Sometimes we’d like our Go programs to intelligently
handle Unix signals. For example, we might want a server to
gracefully shutdown when it receives a SIGTERM, or a
command-line tool to stop processing input if it receives a
SIGINT. Here’s how to handle signals in Go with channels.
*/
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go signal notification works by sending os.Signal values on a
	// channel. We’ll create a channel to receive these notifications
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive notifications
	// of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for signals.
	//When it gets one it’ll print it out and then notify the program
	//that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
