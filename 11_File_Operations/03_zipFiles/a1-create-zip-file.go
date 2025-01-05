package main

import (
	"archive/zip"
	"fmt"
	"os"
)

/*
creating files, adding content and creating zip file
*/
func main() {
	zipFile, err := os.Create("example.zip")
	if err != nil {
		fmt.Println("Error creating ZIP file:", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		fileWriter, err := zipWriter.Create(file)
		if err != nil {
			fmt.Println("Error creating file in ZIP:", err)
			return
		}
		_, err = fileWriter.Write([]byte("Hello, " + file))
		if err != nil {
			fmt.Println("Error writing to file in ZIP:", err)
			return
		}
	}

	fmt.Println("ZIP file created successfully!")
}

// Assignment: Enhance this example, by creating these files and
// adding content in init() function, and create zip file

// Assignment: Using
// file, err := os.Open
// file.Read
// read all .go files in current directory, and create a zip file
