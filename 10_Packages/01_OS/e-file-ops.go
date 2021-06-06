package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetCurrWorkDir() string {
	currDir, err := os.Getwd()
	checkerr(err)
	return currDir
}

func main() {
	fmt.Println("Current Directory:", GetCurrWorkDir())

	// os.Chdir("../")
	// os.Chdir("../")
	os.Chdir("../../")
	fmt.Println("Current Directory:", GetCurrWorkDir())

	// Reads curr direct files
	files, err := os.ReadDir("./") // ioutil.ReadDir(".")
	checkerr(err)
	for _, file := range files {
		fmt.Println(file.IsDir(), file.Name())
	}

	err = filepath.Walk("./", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	checkerr(err)
	fmt.Println(" ", p, info.IsDir())
	return nil
}

// Assignment - Get count of file types in your project directory
// HINT:  map[string][]string{}
//             extn     filename1, filename2
