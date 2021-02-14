package main

import (
	"fmt"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Reading content from a file
	content, err := ioutil.ReadFile("myFile.txt")
	check(err)
	fmt.Printf("File contents: %s\n", content)

	// Writing to a new file newFile.txt.
	data := []byte("Hello, World!")
	err = ioutil.WriteFile("newFile.txt", data, 0644)
	check(err)

	// Printing all files in the current directory.
	// Notice a new newFile.txt file (that we created above).
	files, err := ioutil.ReadDir(".")
	check(err)
	fmt.Println("Files in the current directory:")
	for _, file := range files {
		fmt.Println("\t", file.Name())
	}

}
