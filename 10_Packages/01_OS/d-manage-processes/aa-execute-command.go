package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
	fmt.Println("Process started running ...")
}
