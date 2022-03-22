package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Create our Temp File:  This will create a filename like /tmp/prefix-123456
	// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	// Remember to clean up the file afterwards
	defer os.Remove(tmpFile.Name())

	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	text := []byte("This is a golangcode.com example!")
	if _, err = tmpFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
}
