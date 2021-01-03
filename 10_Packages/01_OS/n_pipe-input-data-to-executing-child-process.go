package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")

	stdin, err := cmd.StdinPipe()

	stdin.Write([]byte("Hello World!")) // <------ here
	stdin.Close()

	if err != nil {
		panic(err)
	}

	data, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	for k, v := range data {
		fmt.Printf("key :  %v, value :  %v \n", k, string(v))
	}
}
