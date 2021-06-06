package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get temporary directory path
	tempDir := os.TempDir()
	fmt.Println("Temp Dir:", tempDir)

	
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
