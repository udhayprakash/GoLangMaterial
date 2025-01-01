package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Try pressing CTRL + C to handle the interrupt")

	select {
	case <-sigChannel:
		fmt.Println("The interrupt got handled")
		os.Exit(0)
	// default:      // Absence of default leads to infinite loop
	// 	fmt.Println("default")
	}
}
