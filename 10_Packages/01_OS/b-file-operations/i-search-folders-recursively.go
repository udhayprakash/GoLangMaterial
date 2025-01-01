package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Function to find and count files with specific extensions in a directory tree
func findFilesByExtensions(root string, extensions []string) (map[string]int, error) {
	// Map to store the count of files for each extension
	fileCounts := make(map[string]int)

	// Walk the directory tree
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		// Check if the file has one of the desired extensions
		if !info.IsDir() {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					fileCounts[ext]++
					break // Stop checking other extensions once a match is found
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the directory tree: %v", err)
	}

	return fileCounts, nil
}

func main() {
	// Extensions to look for
	extensions := []string{".go", ".sh", ".py"} // Add or remove extensions as needed

	// Directory to search
	root := "./"

	// Call the function to find and count files
	fileCounts, err := findFilesByExtensions(root, extensions)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the results
	fmt.Println("File counts by extension:")
	for ext, count := range fileCounts {
		fmt.Printf("%s: %d\n", ext, count)
	}
}
