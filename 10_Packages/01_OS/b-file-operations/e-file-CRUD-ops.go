package main

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// File Creation Example
	file, err := os.Create("example.txt")
	CheckError(err)
	defer file.Close()
	fmt.Println("File 'example.txt' created.")

	// File Writing Example
	content := []byte("Hello, Go file operations!\n")
	err = os.WriteFile("example.txt", content, 0644)
	CheckError(err)
	fmt.Println("Content written to 'example.txt'.")

	// Working with File Stat
	info, err := os.Stat("example.txt")
	if err == nil {
		fmt.Printf("File Name: %s\n", info.Name())
		fmt.Printf("File Size: %d bytes\n", info.Size())
		fmt.Printf("Is Directory: %v\n", info.IsDir())
	} else {
		fmt.Println("File does not exist.")
	}

	// File Deletion Example
	err = os.Remove("example.txt")
	CheckError(err)
	fmt.Println("'example.txt' deleted.")

}
