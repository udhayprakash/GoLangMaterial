package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("binary=", binary)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

//C:\Program Files (x86)\Gow\bin\ls.exe ls -a -l -h
//ls -a -l -h
