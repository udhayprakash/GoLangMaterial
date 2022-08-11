package main

import (
	"fmt"
	"os"
)

func CheckError(er error) {
	if er != nil {
		fmt.Println("Error:", er)
	}
}

func main() {
	// To create a new directory
	output := os.Mkdir("MyDIR", 755)
	fmt.Println("output=", output)

	dirInfo, err := os.Stat("MyDIR")
	CheckError(err)
	fmt.Println("dirInfo:", dirInfo)

	cwd, err := os.Getwd()
	CheckError(err)
	fmt.Println("current working dir:", cwd)

	// changing the current working directory
	os.Chdir("MyDIR")
	cwd, err = os.Getwd()
	CheckError(err)
	fmt.Println("current working dir:", cwd)
}
