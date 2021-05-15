package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	listFilesDirs := exec.Command("ls")
	listOfFilesDirs, err := listFilesDirs.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("listOfFilesDirs=", string(listOfFilesDirs))

	//listFilesDirs1 := exec.Command("ls -ltr") // panic: exec: "ls -ltr": executable file not found in %PATH%

	listFilesDirs1 := exec.Command("ls", "-ltr")
	listOfFilesDirs1, err := listFilesDirs1.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("listOfFilesDirs1=", string(listOfFilesDirs1))

	// To ping google.com and display result
	pingCmd := exec.Command("ping", "google.com")
	pingOutput, pingErr := pingCmd.Output()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("pingOutput =", string(pingOutput))

	//
	output, err := exec.Command("cmd", "/c", "node -v").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

}
