package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("This is first line")
	file.WriteString(`
		This is second line 
		This is third line
	`)
}