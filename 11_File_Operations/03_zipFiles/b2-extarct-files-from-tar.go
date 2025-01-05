package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func main() {
	tarFile, err := os.Open("example.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	tarReader := tar.NewReader(tarFile)

	// Extract each file from the TAR archive
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err.Error() == "EOF" {
				break // End of archive
			}
			log.Fatal(err)
		}

		extractedFile, err := os.Create(header.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer extractedFile.Close()

		// Copy the contents of the file from the TAR archive to the extracted file
		if _, err := io.Copy(extractedFile, tarReader); err != nil {
			log.Fatal(err)
		}

		log.Printf("Extracted: %s\n", header.Name)
	}

	log.Println("Files extracted successfully!")
}
