package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("os.PathSeparator:%v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator:%c\n", os.PathSeparator)

	fmt.Printf("os.DevNull       :%v\n", os.DevNull)

	// To get the current working directory
	fmt.Println(os.Getwd())
	// /workspaces/GolangBatchDec2024/10_Packages/01_OS/b-file-operations <nil>

	cwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)
	fmt.Println("current working directory:", cwd)

	// List files in current directory
	files, err := os.ReadDir("./")
	checkError(err)
	fmt.Printf("files - Type : %T\n", files) // []fs.DirEntry

	for _, file := range files {
		fmt.Printf("file.Name()  : %v\n", file.Name())
		fmt.Printf("file.IsDir() : %v\n\n", file.IsDir())

	}
	fmt.Println()

	// List files in Parent directory
	files, err = os.ReadDir("../")
	checkError(err)
	fmt.Printf("files - Type : %T\n", files) // []fs.DirEntry

	for _, file := range files {
		fmt.Printf("file.Name()  : %v\n", file.Name())
		fmt.Printf("file.IsDir() : %v\n\n", file.IsDir())

	}
}

// /workspaces/GolangBatchDec2024/10_Packages/01_OS/b-file-operations/a_file_paths.go
// \workspaces\GolangBatchDec2024\10_Packages\01_OS\b-file-operations\a_file_paths.go

// my\newFolder

// ioutil.ReadDir("./") also gives same result.
// But this module is getting deprecated
