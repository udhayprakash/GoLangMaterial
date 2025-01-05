package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the GZIP file
	gzipFile, err := os.Open("example.gz")
	if err != nil {
		fmt.Println("Error opening GZIP file:", err)
		return
	}
	defer gzipFile.Close()

	// Create a new GZIP reader
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println("Error creating GZIP reader:", err)
		return
	}
	defer gzipReader.Close()

	// Create the decompressed file
	extractedFile, err := os.Create("extracted.txt")
	if err != nil {
		fmt.Println("Error creating decompressed file:", err)
		return
	}
	defer extractedFile.Close()

	// Copy the contents of the GZIP file to the decompressed file
	if _, err := io.Copy(extractedFile, gzipReader); err != nil {
		fmt.Println("Error copying GZIP contents:", err)
		return
	}

	fmt.Println("File decompressed successfully!")
}
