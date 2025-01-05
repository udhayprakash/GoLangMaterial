package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Open the ZIP file
	zipReader, err := zip.OpenReader("example.zip")
	if err != nil {
		fmt.Println("Error opening ZIP file:", err)
		return
	}
	defer zipReader.Close()

	// Extract each file from the ZIP archive
	for _, file := range zipReader.File {
		// Open the file inside the ZIP archive
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file in ZIP:", err)
			return
		}
		defer fileReader.Close()

		// Create the extracted file
		extractedFile, err := os.Create(filepath.Base(file.Name))
		if err != nil {
			fmt.Println("Error creating extracted file:", err)
			return
		}
		defer extractedFile.Close()

		// Copy the contents of the file from the ZIP archive to the extracted file
		_, err = io.Copy(extractedFile, fileReader)
		if err != nil {
			fmt.Println("Error copying file contents:", err)
			return
		}

		fmt.Printf("Extracted: %s\n", file.Name)
	}

	fmt.Println("Files extracted successfully!")
}
