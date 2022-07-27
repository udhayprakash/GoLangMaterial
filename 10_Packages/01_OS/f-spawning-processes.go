package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// Get Executable file path
	exeDir, _ := os.Executable()
	fmt.Println("Exec Dir:", exeDir)

	fmt.Println("\nruntime.GOOS:", runtime.GOOS)
	command := "ls"
	// if runtime.GOOS == "windows" {
	// 	command = "dir"
	// }
	fmt.Println("Command to run:", command)

	filesNDirs, err := exec.Command(command).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(filesNDirs))

	// filesNDirs, err = exec.Command("ls -ltr").Output() // panic: exec: "ls -ltr": executable file not found in %PATH%
	filesNDirs, err = exec.Command("ls", "-ltr").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(filesNDirs))

	// To ping google.com and display result
	pingCmd := exec.Command("ping", "google.com")
	pingOutput, pingErr := pingCmd.Output()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("pingOutput =", string(pingOutput))

	output, err := exec.Command("cmd", "/c", "node -v").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}


/*
package main

import (
    "log"
    "os/exec"
)

func main() {
    output, err := exec.Command("chrome", "--allow-insecure-localhost").Output()
    if err != nil {
        log.Fatal(err)
    }

fmt.Println(string(out))
}

*/