package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	zipReader, err := zip.OpenReader("example.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Extract each file from the ZIP archive
	for _, file := range zipReader.File {
		fileReader, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer fileReader.Close()

		// Create the extracted file
		extractedFile, err := os.Create(filepath.Base(file.Name))
		if err != nil {
			log.Fatal(err)
		}
		defer extractedFile.Close()

		// Copy the contents of the file from the ZIP archive to the extracted file
		_, err = io.Copy(extractedFile, fileReader)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Extracted: %s\n", file.Name)
	}

	log.Println("Files extracted successfully!")
}
