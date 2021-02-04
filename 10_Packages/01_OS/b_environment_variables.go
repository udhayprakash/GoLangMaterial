package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Retrieving the environment variables
	fmt.Println("os.Environ():", os.Environ())

	for _, each := range os.Environ() {
		//fmt.Printf("each : %T %v\n", each, each)
		//pair := strings.Split(each, "=")
		pair := strings.SplitN(each, "=", 2)
		fmt.Println(pair[0])
	}

	// Setting the environment variables
	os.Setenv("langauge", "Golang")
	os.Setenv("version", "1.15")
	os.Setenv("isBuild", "false")

	// Get env values
	fmt.Println(os.Getenv("langauge"))

	// if the env variable is not set, it will result in empty string
	language123 := os.Getenv("langauge123")
	fmt.Println("language123 =", language123)

	// retrieving builtin environment variables
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Printf(`
	home  : %s
	user  : %s
	gopath:%s`, home, user, gopath)

}

// place all the environment variables in an array and display at the end
