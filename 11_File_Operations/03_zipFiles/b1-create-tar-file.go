package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	tarFile, err := os.Create("example.tar")
	if err != nil {
		fmt.Println("Error creating TAR file:", err)
		return
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		fileContent, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Create a new TAR header
		header := &tar.Header{
			Name: file,
			Mode: 0600,
			Size: int64(len(fileContent)),
		}

		// Write the header to the TAR archive
		if err := tarWriter.WriteHeader(header); err != nil {
			fmt.Println("Error writing TAR header:", err)
			return
		}

		// Write the file content to the TAR archive
		if _, err := tarWriter.Write(fileContent); err != nil {
			fmt.Println("Error writing file content to TAR:", err)
			return
		}
	}

	fmt.Println("TAR file created successfully!")
}
