package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Retrieving the environment variables
	// fmt.Println("os.Environ():", os.Environ())

	envNames := []string{}
	for _, each := range os.Environ() {
		// fmt.Printf("each : %T %v\n", each, each)
		pair := strings.Split(each, "=")
		// fmt.Println(pair[0])
		envNames = append(envNames, pair[0])
	}
	fmt.Println("ENV Names in your system:", envNames)

	// Setting the environment variables
	os.Setenv("language", "Golang")
	os.Setenv("version", "1.15")
	os.Setenv("isBuild", "false")

	// Get env values
	fmt.Println(os.Getenv("language"))

	// if the env variable is not set, it will result in empty string
	language123 := os.Getenv("language123")
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

	os.Setenv("TRAINER", "ravi")

	fmt.Println(os.ExpandEnv(`
	home  : $HOME
	user  : $USER
	gopath: $GOPATH
	Trainer: $TRAINER`))

	path := os.Getenv("PATH")
	gobin := filepath.Join(gopath, "bin")
	if !strings.Contains(path, gobin) {
		log.Fatalf("Your PATH does not contain %s", gobin)
	}
	fmt.Println("gobin:", gobin)

	value, isPresent := os.LookupEnv("SOMETHING")
	fmt.Println(value, isPresent) // "" false

	value, isPresent = os.LookupEnv("GOPATH")
	fmt.Println(value, isPresent) // C:\Users\Amma\go true

}
