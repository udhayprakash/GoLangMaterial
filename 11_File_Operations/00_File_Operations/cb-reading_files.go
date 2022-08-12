package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("myFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Convert []byte to string and print to screen
	text := string(fileContent)
	fmt.Println(text)
}
