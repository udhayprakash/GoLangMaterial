package main

import (
	"os/exec"
)

func myfunc() (string, error) {

	grep := exec.Command("grep", "func", "pipe.go")
	out, err := grep.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := grep.Start(); err != nil {
		return "", err
	}
	wc := exec.Command("wc", "-l")
	wc.Stdin = out
	data, err := wc.CombinedOutput()
	return string(data), nil
}
