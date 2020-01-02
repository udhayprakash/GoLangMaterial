package main

import (
	"fmt"
	"log"
	"regexp"
)
/*
Purpose: To remove all non alphanumerical characters from a string
*/
func main() {

	example := "#GoLangCode!$!"

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf("A string of %s becomes %s \n", example, processedString)
}