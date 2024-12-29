package main

import (
	"fmt"
)

type UploadError struct {
	message string
	file    string
}

func (e UploadError) Error() string {
	return fmt.Sprintf("Upload failed for file '%s': %v", e.file, e.message)
}

func uploadFile(fileName string) error {
	return UploadError{"file too large", fileName}
}

func main() {
	if err := uploadFile("sample.png"); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
