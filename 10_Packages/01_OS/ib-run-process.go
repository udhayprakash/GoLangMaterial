package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("calc")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
