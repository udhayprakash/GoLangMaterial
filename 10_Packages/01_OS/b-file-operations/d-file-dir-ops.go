package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileExits(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {

	// To get current working directory
	currDir, err := os.Getwd()
	check(err)
	fmt.Println("Current Directory:", currDir)

	// To create a sub-directory
	// Cleanup - remove all if exits
	if FileExits("subdir") == true {
		os.RemoveAll("subdir") // rm -rf
	}

	// To create a single directory
	err = os.Mkdir("subdir", 0755)
	check(err)

	// Method 1-  To create a flat file
	err = ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	check(err)

	// To read and display all files & dirs
	files, err := os.ReadDir("./")
	check(err)
	for _, file := range files {
		fmt.Println(file.IsDir(), file.Name())
	}

	// Method 2- To create a flat file
	fh, err := os.Create("newfile.txt")
	check(err)

	fh.WriteString("This is first line")
	fh.Write([]byte{115, 111, 109, 101, 10})

	// creating a writer object
	wr := bufio.NewWriter(fh)
	wr.WriteString("Hello world\n")
	wr.WriteString("Second Line\n")
	wr.WriteString("I am a developer\n")
	wr.Write([]byte{115, 111, 109, 101, 10})
	wr.Flush()

	// To create one or more directories in chain
	err = os.MkdirAll("otherDir", 0700)
	check(err)

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// To read and display all files & dirs
	files, err = os.ReadDir("./")
	check(err)
	for _, file := range files {
		fmt.Println(file.IsDir(), file.Name())
	}
}
