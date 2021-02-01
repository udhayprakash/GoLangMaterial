package main

import (
	"fmt"
	"os"
)

func main() {
	// open xml file
	xmlFile, err := os.Open("users.xml")
	defer xmlFile.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully opened users.xml")
	}
}
