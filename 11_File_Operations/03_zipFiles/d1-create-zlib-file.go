package main

import (
	"compress/zlib"
	"fmt"
	"os"
)

func main() {
	// Create a new ZLIB file
	zlibFile, err := os.Create("example.zlib")
	if err != nil {
		fmt.Println("Error creating ZLIB file:", err)
		return
	}
	defer zlibFile.Close()

	// Create a new ZLIB writer
	zlibWriter := zlib.NewWriter(zlibFile)
	defer zlibWriter.Close()

	// Write content to the ZLIB file
	content := []byte("Hello, ZLIB!")
	if _, err := zlibWriter.Write(content); err != nil {
		fmt.Println("Error writing to ZLIB file:", err)
		return
	}

	fmt.Println("ZLIB file created successfully!")
}
