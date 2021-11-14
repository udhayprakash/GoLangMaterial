package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		.   	current directory
		.. 		parent directory
		../.. 	parent's parent directory

	*/
	dir, err := os.Open("../..")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
