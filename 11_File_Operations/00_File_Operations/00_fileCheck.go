package main

import (
	"fmt"
	"os"
)

func main() {
	for _, eachFile := range [2]string{"example.txt", "C:\\Go\\README.md"} {
		if fileExists(eachFile) {
			fmt.Printf("file - %s exists\n", eachFile)
		} else {
			fmt.Printf("file - %s does not exist (or is a directory)\n", eachFile)
		}
	}

}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
