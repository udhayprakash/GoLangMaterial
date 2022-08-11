package main

import (
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	fmt.Println("cwd:", cwd)

	//file, err := os.Open("10_File_Operations\\01_File_Operations\\myFile.txt")
	file, err := os.Open("myFile.txt")
	if err != nil {
		// handle the error here
		fmt.Println("err:", err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println("stat:", stat)

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	//fmt.Println("bs:", bs)

	fileContent := string(bs)
	fmt.Println("fileContent:", fileContent)

	// filename, _ := filepath.Abs("./fruits.yml")
}
