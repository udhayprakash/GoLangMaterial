package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	// Create a new GZIP file
	gzipFile, err := os.Create("example.gz")
	if err != nil {
		fmt.Println("Error creating GZIP file:", err)
		return
	}
	defer gzipFile.Close()

	// Create a new GZIP writer
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	// Write content to the GZIP file
	content := []byte("Hello, GZIP!")
	if _, err := gzipWriter.Write(content); err != nil {
		fmt.Println("Error writing to GZIP file:", err)
		return
	}

	fmt.Println("GZIP file created successfully!")
}
