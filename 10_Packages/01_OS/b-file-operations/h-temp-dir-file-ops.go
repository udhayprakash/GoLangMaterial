package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. Get the system's default temporary directory
	tempDir := os.TempDir()
	fmt.Printf("System's default temporary directory: %s\n", tempDir)

	// 2. Create a new temporary directory
	customTempDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return
	}
	fmt.Printf("Created temporary directory: %s\n", customTempDir)

	// 3. Create a temporary file in the custom temporary directory
	// tempFile, err := ioutil.TempFile(customTempDir, "tempfile-*.txt")
	tempFile, err := os.CreateTemp(customTempDir, "tempfile-*.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	// 4. Write some data to the temporary file
	data := []byte("Hello, this is a temporary file!")
	if _, err := tempFile.Write(data); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
	fmt.Println("Data written to temporary file.")

	// 5. Read the data back from the temporary file
	if _, err := tempFile.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	// readData, err := ioutil.ReadAll(tempFile)
	readData, err := io.ReadAll(tempFile)
	if err != nil {
		fmt.Println("Error reading from temporary file:", err)
		return
	}
	fmt.Printf("Data read from temporary file: %s\n", string(readData))

	// 6. Clean up: Close the file and remove the temporary directory and its contents
	if err := tempFile.Close(); err != nil {
		fmt.Println("Error closing temporary file:", err)
	}
	if err := os.RemoveAll(customTempDir); err != nil {
		fmt.Println("Error removing temporary directory:", err)
	}
	fmt.Println("Temporary directory and file cleaned up.")
}
