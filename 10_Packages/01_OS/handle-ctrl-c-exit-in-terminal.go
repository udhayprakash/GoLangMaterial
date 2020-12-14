package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const FileNameExample = "go-example.txt"

func main() {

	// Setup our Ctrl+C handler
	SetupCloseHandler()

	// Run our program... We create a file to clean up then sleep
	CreateFile()
	for {
		fmt.Println("- Sleeping")
		time.Sleep(10 * time.Second)
	}
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		DeleteFiles()
		os.Exit(0)
	}()
}

// DeleteFiles is used to simulate a 'clean up' function to run on shutdown. Because
// it's just an example it doesn't have any error handling.
func DeleteFiles() {
	fmt.Println("- Run Clean Up - Delete Our Example File")
	_ = os.Remove(FileNameExample)
	fmt.Println("- Good bye!")
}

// Create a file so we have something to clean up when we close our program.
func CreateFile() {
	fmt.Println("- Create Our Example File")
	file, _ := os.Create(FileNameExample)
	defer file.Close()
}
