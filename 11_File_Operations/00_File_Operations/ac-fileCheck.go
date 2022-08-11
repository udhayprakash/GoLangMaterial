package main

import (
	"fmt"
	"io/fs"
	"os"
)

func fileExists(filename string) (bool, fs.FileInfo) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	}
	fmt.Println("\n\tinfo.IsDir():", info.IsDir())
	return info != nil, info
}

func main() {
	//// if the file is not present
	//info, err := os.Stat("example.txt")
	//fmt.Println("info:", info)
	//fmt.Println("err :", err)

	//// when the file is present
	//info, err = os.Stat("C:\\Go\\README.md")
	//fmt.Println("\ninfo:", info)
	//fmt.Println("err :", err)

	for _, eachFile := range [3]string{
		"example.tx",
		"C:\\Go\\README.md",
		"01_fileCheck.go",
	} {
		//fmt.Println(eachFile, fileExists(eachFile))
		isFilePresent, info := fileExists(eachFile)
		if isFilePresent {
			fmt.Printf("\nfile - %s exists\n", eachFile)
			fmt.Println("\t Modification Time: ", info.ModTime())
			fmt.Println("\t File Mode        : ", info.Mode())
			fmode := info.Mode()
			if fmode.IsRegular() {
				fmt.Println("\tThis is a regular file")
			}
			fmt.Println("\tFile Size         :", info.Size())
			fmt.Println("\tIs Dir            :", info.IsDir())

		} else {
			fmt.Printf("\nfile - %s does not exist (or is a directory)\n", eachFile)
		}
	}

}
