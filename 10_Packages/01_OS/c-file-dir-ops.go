package main

import (
	"fmt"
	"os"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("os.PathSeparator:%v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator:%c\n", os.PathSeparator)

	fmt.Printf("os.DevNull       :%v\n", os.DevNull)



	// To get the current working directory
	cwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	CheckError(err)
	fmt.Println("current working directory:", cwd)

	files, err := os.ReadDir("./")
	CheckError(err)
	fmt.Printf("files - Type : %T\n", files) // []fs.DirEntry

	for _, file := range files {
		fmt.Printf("file.Name()  : %v\n", file.Name())
		fmt.Printf("file.IsDir() : %v\n\n", file.IsDir())

	}
}

// ioutil.ReadDir("./") also gives same result.
// But this module is getting deprecated
