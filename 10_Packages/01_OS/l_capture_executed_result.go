package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	cmd := exec.Command("whoami")

	// capture the output and error pipes
	stdout, err := cmd.StdoutPipe()
	handleError(err)

	err = cmd.Start()
	handleError(err)

	defer cmd.Wait()
	buff := bufio.NewScanner(stdout)

	var returnText []string

	for buff.Scan() {
		returnText = append(returnText, buff.Text())
	}

	fmt.Println(returnText)
}
