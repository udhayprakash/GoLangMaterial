package main

import (
	"fmt"
	"runtime"
)

func main() {

	if runtime.GOOS == "windows" {
		fmt.Println("You are running on Windows")
	} else {
		fmt.Println("You are running on an OS other than Windows")
	}
	fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS)
	fmt.Println("runtime.NumCPU()  :", runtime.NumCPU())
	fmt.Println("runtime.Version() :", runtime.Version())
}
