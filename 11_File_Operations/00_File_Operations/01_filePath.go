package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Make a cross-platform file
	// unix='dir/example' windows='dir\example'

	// Option 1
	examplePath1 := "dir" + string(os.PathSeparator) + "example"
	fmt.Println("PathSeparator: " + examplePath1)

	// Option 2
	examplePath2 := filepath.FromSlash("dir/example")
	fmt.Println("FromSlash: " + examplePath2)
}
