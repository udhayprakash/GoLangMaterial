package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 1. Current Working Directory
	cwd, err := os.Getwd()
	CheckError(err)
	fmt.Println("Current Working Directory:", cwd)

	// 2. Constructing a Path
	fileName := "example.txt"
	fullPath := filepath.Join(cwd, fileName)
	fmt.Println("Constructed Path:", fullPath)

	// 3. Getting Absolute Path
	absPath, err := filepath.Abs(fileName)
	CheckError(err)
	fmt.Println("Absolute Path:", absPath)

	// 4. Getting Base and Directory
	fmt.Println("Base of Path:", filepath.Base(fullPath))          // example.txt
	fmt.Println("Directory of Path:", filepath.Dir(fullPath))     // cwd
	fmt.Println("Extension of Path:", filepath.Ext(fullPath))     // .txt

	// 5. Cleaning a Path
	uncleanPath := "/tmp/../var/log//"
	cleanPath := filepath.Clean(uncleanPath)
	fmt.Println("Cleaned Path:", cleanPath) // /var/log

	// 6. Checking if Path is Absolute
	isAbs := filepath.IsAbs(fullPath)
	fmt.Printf("Is '%s' Absolute? %v\n", fullPath, isAbs)

	// 7. Splitting a Path
	dir, file := filepath.Split(fullPath)
	fmt.Println("Directory Part:", dir)
	fmt.Println("File Part:", file)

	// 8. Walking a Directory Tree
	fmt.Println("Walking through files in current directory:")
	err = filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		CheckError(err)
		fmt.Printf("Visited: %s (IsDir: %v)\n", path, d.IsDir())
		return nil
	})
	CheckError(err)

	// 9. Converting Relative to Absolute
	relativePath := "../"
	convertedPath, err := filepath.Abs(relativePath)
	CheckError(err)
	fmt.Println("Converted to Absolute:", convertedPath)
	// /workspaces/GolangBatchDec2024/10_Packages/01_OS
	
	// 10. Temporary Files and Paths
	tempFile := filepath.Join(os.TempDir(), "tempfile.txt")
	fmt.Println("Temporary File Path:", tempFile)
}