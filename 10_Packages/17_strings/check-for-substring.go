package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "this is a string containing a big string and small string"

	subStr := "big string"

	if strings.Contains(str, subStr) {
		fmt.Printf("Found subStr in str \n")
	} else {
		fmt.Printf("subStr is not in str \n")
	}

	subStr2 := "another string"

	if strings.Contains(str, subStr2) {
		fmt.Printf("Found subStr in str \n")
	} else {
		fmt.Printf("subStr2 is not in str \n")
	}

}
