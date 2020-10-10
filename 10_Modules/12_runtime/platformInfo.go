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

}
