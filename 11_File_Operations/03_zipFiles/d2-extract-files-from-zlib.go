package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the ZLIB file
	zlibFile, err := os.Open("example.zlib")
	if err != nil {
		fmt.Println("Error opening ZLIB file:", err)
		return
	}
	defer zlibFile.Close()

	// Create a new ZLIB reader
	zlibReader, err := zlib.NewReader(zlibFile)
	if err != nil {
		fmt.Println("Error creating ZLIB reader:", err)
		return
	}
	defer zlibReader.Close()

	// Create the decompressed file
	extractedFile, err := os.Create("extracted.txt")
	if err != nil {
		fmt.Println("Error creating decompressed file:", err)
		return
	}
	defer extractedFile.Close()

	// Copy the contents of the ZLIB file to the decompressed file
	if _, err := io.Copy(extractedFile, zlibReader); err != nil {
		fmt.Println("Error copying ZLIB contents:", err)
		return
	}

	fmt.Println("File decompressed successfully!")
}
