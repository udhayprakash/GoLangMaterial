package main

import (
	"fmt"
	"os"
)

func main() {
	//// if the file is not present
	//info, err := os.Stat("example.txt")
	//fmt.Println("info:", info)
	//fmt.Println("err :", err)
	//
	//// when the file is present
	//info, err = os.Stat("C:\\Go\\README.md")
	//fmt.Println("\ninfo:", info)
	//fmt.Println("err :", err)


	for _, eachFile := range [3]string{"example.tx", "C:\\Go\\README.md", "01_File_Operations"}{
		//fmt.Println(eachFile, fileExists(eachFile))
		if fileExists(eachFile) {
			fmt.Printf("file - %s exists\n", eachFile)
		} else {
			fmt.Printf("file - %s does not exist (or is a directory)\n", eachFile)
		}
	}


}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("\n\tinfo.IsDir():", info.IsDir())
	return info != nil
}
