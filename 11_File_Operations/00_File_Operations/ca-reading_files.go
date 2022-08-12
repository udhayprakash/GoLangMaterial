package main

import (
	"fmt"
	"os"
)

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cwd, err := os.Getwd()
	fmt.Println("cwd:", cwd)

	//file, err := os.Open("..\\00_File_Operations\\testfile.txt")
	file, err := os.Open("testfile.txt")
	handleErr(err)
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	handleErr(err)
	fmt.Println("stat:", stat)

	// creating buffer, of buffer size length, to store the read content
	bs := make([]byte, stat.Size())

	// Reading the ENTIRE content from file
	_, err = file.Read(bs)
	handleErr(err)
	// fmt.Println("bs:", bs)
	fileContent := string(bs)
	fmt.Println("fileContent:", fileContent)

}
