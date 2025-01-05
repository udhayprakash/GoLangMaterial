package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	// Open the GZIP file
	gzipFile, err := os.Open("example.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer gzipFile.Close()

	// Create a new GZIP reader
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// Create the decompressed file
	extractedFile, err := os.Create("extracted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer extractedFile.Close()

	// Copy the contents of the GZIP file to the decompressed file
	if _, err := io.Copy(extractedFile, gzipReader); err != nil {
		log.Fatal(err)
	}

	log.Println("File decompressed successfully!")
}
