package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// which go (or) where go
	pathInfo, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}
	fmt.Println("pathInfo:", pathInfo)

	ls_binary_path, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("ls_path=", ls_binary_path)

	args := []string{"-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(ls_binary_path, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// ls -a -l -h
// /usr/bin/ls -a -l -h
// C:\Program Files (x86)\Gow\bin\ls.exe ls -a -l -h
