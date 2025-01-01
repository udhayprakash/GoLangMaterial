package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
