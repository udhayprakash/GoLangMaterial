package main

import (
	"fmt"
	"os"
)

func handleErr(er error) {
	if er != nil {
		// handle the error here
		fmt.Println(er)
		return
	}
}

func main() {
	file, err := os.Create("testfile.txt")
	handleErr(err)
	defer file.Close()

	fmt.Println("file name is", file.Name())
	
	// Writing data to file
	file.WriteString("This is first line")
	file.WriteString(`
		This is second line 
		This is third line
	`)

	// writing individual bytes to the file 
	data := []byte("This is fourth line\n")
	file.Write(data)

	data2 := []byte{'a', 'b', 'c', '1', '2', '@', '\n'}
	file.Write(data2)

	// Sync function will force the data to be written out
	file.Sync()
}
