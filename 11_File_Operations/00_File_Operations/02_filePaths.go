package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	// Make a cross-platform file
	//      unix='dir/example/folder1'
	//      windows='dir\example\folder1'

	// Option 1
	examplePath1 := "dir" + string(os.PathSeparator) + "example" + string(os.PathSeparator) + "folder1"
	fmt.Println("PathSeparator: " + examplePath1)

	// Option 2
	examplePath2 := filepath.FromSlash("dir/example/folder1")
	fmt.Println("FromSlash    : " + examplePath2)

}