package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	tarFile, err := os.Open("example.tar")
	if err != nil {
		fmt.Println("Error opening TAR file:", err)
		return
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
			fmt.Println("Error reading TAR header:", err)
			return
		}

		extractedFile, err := os.Create(header.Name)
		if err != nil {
			fmt.Println("Error creating extracted file:", err)
			return
		}
		defer extractedFile.Close()

		// Copy the contents of the file from the TAR archive to the extracted file
		if _, err := io.Copy(extractedFile, tarReader); err != nil {
			fmt.Println("Error copying file contents:", err)
			return
		}

		fmt.Printf("Extracted: %s\n", header.Name)
	}

	fmt.Println("Files extracted successfully!")
}
