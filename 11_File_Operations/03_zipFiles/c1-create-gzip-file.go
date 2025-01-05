package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	// Create a new GZIP file
	gzipFile, err := os.Create("example.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer gzipFile.Close()

	// Create a new GZIP writer
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	// Write content to the GZIP file
	content := []byte("Hello, GZIP!")
	if _, err := gzipWriter.Write(content); err != nil {
		log.Fatal(err)
	}

	log.Println("GZIP file created successfully!")
}
