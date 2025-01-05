package main

import (
	"archive/tar"
	"log"
	"os"
)

func main() {
	tarFile, err := os.Create("example.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		fileContent, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		// Create a new TAR header
		header := &tar.Header{
			Name: file,
			Mode: 0600,
			Size: int64(len(fileContent)),
		}

		// Write the header to the TAR archive
		if err := tarWriter.WriteHeader(header); err != nil {
			log.Fatal(err)
		}

		// Write the file content to the TAR archive
		if _, err := tarWriter.Write(fileContent); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("TAR file created successfully!")
}
