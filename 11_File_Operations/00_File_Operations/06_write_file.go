package main

import (
	"io"
	"log"
	"os"
)

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	err := WriteToFile("result.txt", "Hello Readers of golangcode.com\n")
	if err != nil {
		log.Fatal(err)
	}
}
